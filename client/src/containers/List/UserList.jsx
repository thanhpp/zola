import React from "react";
import EditTableRow from "../../components/table/EditableTableRow";
import { Avatar } from "antd";
import { UserOutlined } from "@ant-design/icons";
import { Link } from "react-router-dom";

const columns = [
	{
		title: "Username",
		dataIndex: "username",
		key: "username",
		render: (text, row) => {
			const { url } = row;
			return (
				<>
					{url ? (
						<Avatar size="small" src={url} />
					) : (
						<Avatar size="small" icon={<UserOutlined />} />
					)}
					{/* eslint-disable-next-line jsx-a11y/anchor-is-valid */}
					<Link to={`${text}`} style={{ marginLeft: 15 }}>
						{text}
					</Link>
				</>
			);
		},
	},
	{
		title: "Status",
		dataIndex: "status",
		key: "status",
	},
	{
		title: "State",
		dataIndex: "state",
		key: "state",
		editable: true,
	},
	{
		title: "Last login",
		dataIndex: "last_login",
		key: "last_login",
	},
];

const options = [
	{ value: 0, text: "Inactive" },
	{ value: 1, text: "Active" },
];

const data = [
	{
		username: "omg",
		key: "omg",
		url: "",
		status: "online",
		state: "Inactive",
		last_login: "6 hours ago",
	},
	{
		username: "omg2",
		key: "omg2",
		url: "https://joeschmoe.io/api/v1/random",
		status: "online",
		state: "Inactive",
		last_login: "6 hours ago",
	},
];

export default function UserList() {
	console.log("User List render");
	return <EditTableRow columnName={columns} data={data} options={options} />;
}
