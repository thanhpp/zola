import React, { useState } from "react";
import EditTableRow from "../../components/table/EditableTableRow";
import { Avatar } from "antd";
import { UserOutlined } from "@ant-design/icons";
import { Link } from "react-router-dom";
import dayjs from "dayjs";
import relativeTime from "dayjs/plugin/relativeTime";
dayjs.extend(relativeTime);

const columns = [
	{
		title: "User ID",
		dataIndex: "user_id",
		key: "user_id",
	},
	{
		title: "Username",
		dataIndex: "username",
		key: "username",
		render: (text, row) => {
			const { avatar } = row;
			return (
				<>
					{avatar ? (
						<Avatar size="small" src={avatar} />
					) : (
						<Avatar size="small" icon={<UserOutlined />} />
					)}
					<Link to={`${text}`} style={{ marginLeft: 15 }}>
						{text}
					</Link>
				</>
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

const convertedData = users.map((user) => {
	return {
		key: user.user_id,
		user_id: user.user_id,
		username: user.username,
		avatar: user.avatar,
		state: isActive(user.is_active),
		lastLogin: dayjs.unix(user.lastLogin).fromNow(),
	};
});

export default function UserList() {
	const [data, setData] = useState(convertedData);
	const handleAdd = (values) => {
		const { phoneNumber, password } = values;
		console.log({
			phoneNumber: phoneNumber,
			password: password,
		});
	};
	const handleDelete = (id) => {
		console.log(id);
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
			data={data}
			handleAdd={handleAdd}
			handleDelete={handleDelete}
			handleEdit={handleEdit}
			options={options}
		/>
	);
}
