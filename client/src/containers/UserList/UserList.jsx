import React from "react";
import EditTableRow from "../../components/table/EditableTableRow";
import { Avatar } from "antd";
import { UserOutlined } from "@ant-design/icons";

const columns = [
	{
		title: "Username",
		dataIndex: "username",
		key: "username",
		render: (text, url) => {
			<>
				{url ? (
					<Avatar size="small" src={url} />
				) : (
					<Avatar size="small" icon={<UserOutlined />} />
				)}
				{/* eslint-disable-next-line jsx-a11y/anchor-is-valid */}
				<a>{text}</a>
			</>;
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

const data = [];

export default function UserList() {
	return <EditTableRow columnName={columns} data={data} options={options} />;
}
