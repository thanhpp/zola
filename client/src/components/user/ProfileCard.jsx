import React from "react";
import "antd/dist/antd.css";
import { Card, Avatar } from "antd";
import {
	EditOutlined,
	EllipsisOutlined,
	SettingOutlined,
} from "@ant-design/icons";

const { Meta } = Card;

export default function ProfileCard() {
	return (
		<Card
			// style={{ width: "30%" }}
			cover={
				<img
					alt="example"
					style={{ objectFit: "cover" }}
					src="https://gw.alipayobjects.com/zos/rmsportal/JiqGstEfoWAOHiTxclqi.png"
				/>
			}
			actions={[
				<SettingOutlined key="setting" />,
				<EditOutlined key="edit" />,
				<EllipsisOutlined key="ellipsis" />,
			]}
		>
			<Meta
				avatar={<Avatar size={64} src="https://joeschmoe.io/api/v1/random" />}
				style={{ display: "flex", justifyContent: "center" }}
			/>
		</Card>
	);
}
