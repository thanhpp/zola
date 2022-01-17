import React from "react";
import EditTableRow from "../../components/table/EditableTableRow";
import { Avatar, Space, message } from "antd";
import { UserOutlined } from "@ant-design/icons";
import { Link } from "react-router-dom";
import { useMutation, useQuery, useQueryClient } from "react-query";
import { signUpUser } from "../../api/userAuthentication";
import { getUserList, deleteUser, setUserState } from "../../api/userApi";

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
		title: "Phone number",
		dataIndex: "phonenumber",
		key: "phonenumber",
	},
	{
		title: "State",
		dataIndex: "state",
		key: "state",
		editable: true,
	},
];

const options = [
	{ value: "locked", text: "Inactive" },
	{ value: "active", text: "Active" },
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
			phonenumber: user.phone,
		};
	});
};

export default function UserList() {
	const queryClient = useQueryClient();
	const { data: query, isLoading } = useQuery("users", getUserList);
	const { mutate: addUserMutation } = useMutation(signUpUser, {
		onSuccess: () => {
			queryClient.invalidateQueries("users");
		},
		onError: (error) => {
			message.error({
				content: `Code: ${error.response?.data?.code};
				Message: ${error.response?.data?.message}`,
			});
		},
	});
	const { mutate: deleteUserMutation } = useMutation(deleteUser, {
		onSuccess: () => {
			queryClient.invalidateQueries("users");
		},
		onError: (error) => {
			message.error({
				content: `Code: ${error.response?.data?.code};
				Message: ${error.response?.data?.message}`,
			});
		},
	});

	const { mutate: setUserStateMutation } = useMutation(setUserState, {
		onSuccess: () => {
			queryClient.invalidateQueries("users");
		},
		onError: (error) => {
			message.error({
				content: `Code: ${error.response?.data?.code};
				Message: ${error.response?.data?.message}`,
			});
		},
		onMutate: () => {
			message.loading("loading");
		},
	});

	const handleAdd = (values) => {
		addUserMutation(values);
	};
	const handleDelete = (id) => {
		deleteUserMutation(id);
	};
	const handleEdit = (values) => {
		setUserStateMutation(values);
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
