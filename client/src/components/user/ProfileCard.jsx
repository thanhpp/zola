import React, { useContext, useState } from "react";
import "antd/dist/antd.css";
import { Card, Avatar, Tooltip, Tag } from "antd";
import {
	UserSwitchOutlined,
	UserDeleteOutlined,
	UserOutlined,
} from "@ant-design/icons";

import AuthContext from "../../context/authContext";

const { Meta } = Card;

export default function ProfileCard({ user }) {
	const authCtx = useContext(AuthContext);
	const { cover_image, avatar, is_active, name, description, is_online, id } =
		user;

	const [userState, setUserState] = useState(!!!is_active);

	const handleDeleteUser = () => {
		console.log("delete user", id);
	};

	const handleUserState = () => {
		//async request here
		console.log(`new state: ${!userState ? "1" : "0"}, userId: ${id}`);
		// handle in client
		setUserState(!userState);
	};

	return (
		<Card
			cover={
				cover_image ? (
					<img
						alt="example"
						style={{ objectFit: "cover", height: "15rem" }}
						src={cover_image}
					/>
				) : (
					<img
						alt="example"
						style={{ objectFit: "cover", height: "15rem" }}
						src={"https://bom.so/puj8BQ"}
					/>
				)
			}
			actions={
				id === authCtx.user.userId
					? []
					: [
							<Tooltip placement="bottom" title="Delete user">
								<UserDeleteOutlined key="delete" onClick={handleDeleteUser} />
							</Tooltip>,
							<Tooltip
								placement="bottom"
								title={userState ? "Set user inactive" : "Set user active"}
							>
								<UserSwitchOutlined key="state" onClick={handleUserState} />
							</Tooltip>,
					  ]
			}
		>
			<Meta
				avatar={
					avatar ? (
						<Avatar size={64} src={avatar} />
					) : (
						<Avatar size={64} icon={<UserOutlined />} />
					)
				}
				title={name}
				description={description}
			/>
			{parseInt(is_online) ? (
				<Tag color={"green"}>Online</Tag>
			) : (
				<Tag color={"default"}>Offline</Tag>
			)}
			{parseInt(is_active) ? (
				<Tag color={"geekblue"}>Active</Tag>
			) : (
				<Tag color={"default"}>Inactive</Tag>
			)}
		</Card>
	);
}
