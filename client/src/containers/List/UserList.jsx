import React, { useState } from "react";
import EditTableRow from "../../components/table/EditableTableRow";
import { Avatar, Space } from "antd";
import { UserOutlined } from "@ant-design/icons";
import { Link } from "react-router-dom";
import dayjs from "dayjs";
import relativeTime from "dayjs/plugin/relativeTime";
import { useMutation, useQuery, useQueryClient } from "react-query";
import { signUpUser } from "../../api/userAuthentication";
import { getUserList, deleteUser } from "../../api/userApi";
dayjs.extend(relativeTime);

const columns = [
	{
		title: "User ID",
		dataIndex: "user_id",
		key: "user_id",
		render: (text) => {
			return <Link to={`${text}`}>{text}</Link>;
		},
	},
	{
		title: "Name",
		dataIndex: "name",
		key: "name",
		render: (text, row) => {
			const { avatar } = row;
			return (
				<Space>
					{avatar ? (
						<Avatar size="small" src={avatar} />
					) : (
						<Avatar size="small" icon={<UserOutlined />} />
					)}

					{text}
				</Space>
			);
		},
	},
	{
		title: "State",
		dataIndex: "state",
		key: "state",
		editable: true,
	},
	{
		title: "Last login",
		dataIndex: "lastLogin",
		key: "lastLogin",
	},
];

const options = [
	{ value: 0, text: "Inactive" },
	{ value: 1, text: "Active" },
];

const users = [
	{
		user_id: "omg",
		username: "omg",
		avatar: "https://joeschmoe.io/api/v1/random",
		is_active: "0",
		lastLogin: "1639121121",
	},
	{
		user_id: "omg 2",
		username: "omg",
		avatar: "https://joeschmoe.io/api/v1/random",
		is_active: "1",
		lastLogin: "1639121111",
	},
];

const isActive = (state) => {
	return parseInt(state) ? "Active" : "Inactive";
};

const convertedData = (query) => {
	if (!query) return;
	return query.data.users.map((user) => {
		return {
			key: user.user_id,
			user_id: user.user_id,
			name: user.name,
			avatar: user.avatar,
			state: isActive(user.is_active),
			lastLogin: !user.lastLogin ? "" : dayjs.unix(user.lastLogin).fromNow(),
		};
	});
};

export default function UserList() {
	const queryClient = useQueryClient();
	const [data, setData] = useState(users);
	const { data: query, isLoading } = useQuery("users", getUserList);
	//console.log(query);
	const { mutate: addUserMutation } = useMutation(signUpUser, {
		onSuccess: (data) => {
			console.log("added user", data);
			queryClient.invalidateQueries("users");
		},
	});
	const { mutate: deleteUserMutation } = useMutation(deleteUser, {
		onSuccess: (data) => {
			console.log("deleted", data);
			queryClient.invalidateQueries("users");
		},
	});
	const handleAdd = (values) => {
		addUserMutation(values);
	};
	const handleDelete = (id) => {
		console.log(id);
		deleteUserMutation(id);
	};
	const handleEdit = (values) => {
		//edit user state - send async request
		const { user_id, state } = values;
		console.log(user_id, state);
		//handle array client
		values = { ...values, state: isActive(state) };
		const newData = data.map((data) => {
			if (data.user_id === values.user_id) return values;
			return data;
		});
		setData(newData);
	};
	return (
		<EditTableRow
			columnName={columns}
			loading={isLoading}
			data={convertedData(query)}
			handleAdd={handleAdd}
			handleDelete={handleDelete}
			handleEdit={handleEdit}
			options={options}
		/>
	);
}
