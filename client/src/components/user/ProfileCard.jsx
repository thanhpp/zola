import React, { useState } from "react";
import "antd/dist/antd.css";
import { Card, Avatar, Tooltip, Tag } from "antd";
import { UserSwitchOutlined, UserDeleteOutlined } from "@ant-design/icons";

const { Meta } = Card;

export default function ProfileCard(props) {
	const { coverImg, avatar, state, username, description, online, userId } =
		props;

	const [userState, setUserState] = useState(!!!state);

	const handleDeleteUser = () => {
		console.log("delete user", userId);
	};

	const handleUserState = () => {
		//async request here
		console.log(`new state: ${!userState ? "1" : "0"}, userId: ${userId}`);
		// handle in client
		setUserState(!userState);
	};

	return (
		<Card
			cover={
				<img alt="example" style={{ objectFit: "cover" }} src={coverImg} />
			}
			actions={[
				<Tooltip placement="bottom" title="Delete user">
					<UserDeleteOutlined key="delete" onClick={handleDeleteUser} />
				</Tooltip>,
				<Tooltip
					placement="bottom"
					title={userState ? "Set user inactive" : "Set user active"}
				>
					<UserSwitchOutlined key="state" onClick={handleUserState} />
				</Tooltip>,
			]}
		>
			<Meta
				avatar={<Avatar size={64} src={avatar} />}
				title={username}
				description={description}
			/>
			{parseInt(online) ? (
				<Tag color={"green"}>Online</Tag>
			) : (
				<Tag color={"default"}>Offline</Tag>
			)}
			{userState ? (
				<Tag color={"geekblue"}>Active</Tag>
			) : (
				<Tag color={"default"}>Inactive</Tag>
			)}
		</Card>
	);
}
