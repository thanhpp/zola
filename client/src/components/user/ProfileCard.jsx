import React from "react";
import "antd/dist/antd.css";
import { Card, Avatar, Tag } from "antd";
import { UserOutlined } from "@ant-design/icons";

const { Meta } = Card;

export default function ProfileCard({ user }) {
	const { cover_image, avatar, is_active, name, description, is_online } = user;

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
