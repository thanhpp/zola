import React from "react";
import "antd/dist/antd.css";
import { Card, Avatar, Tooltip } from "antd";
import { UserSwitchOutlined, UserDeleteOutlined } from "@ant-design/icons";

const { Meta } = Card;

export default function ProfileCard() {
	const state = 0;

	const handleDeleteUser = () => {
		console.log("delete user");
	};

	const handleUserState = () => {
		console.log("user state");
	};

	return (
		<Card
			cover={
				<img
					alt="example"
					style={{ objectFit: "cover" }}
					src="https://gw.alipayobjects.com/zos/rmsportal/JiqGstEfoWAOHiTxclqi.png"
				/>
			}
			actions={[
				<Tooltip placement="bottom" title="Delete user">
					<UserDeleteOutlined key="delete" onClick={handleDeleteUser} />
				</Tooltip>,
				<Tooltip
					placement="bottom"
					title={state ? "Set user inactive" : "Set user active"}
				>
					<UserSwitchOutlined key="state" onClick={handleUserState} />
				</Tooltip>,
			]}
		>
			<Meta
				avatar={<Avatar size={64} src="https://joeschmoe.io/api/v1/random" />}
				style={{ display: "flex", justifyContent: "center" }}
			/>
		</Card>
	);
}
