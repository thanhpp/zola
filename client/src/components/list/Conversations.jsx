import React from "react";
import "antd/dist/antd.css";
import { List, Avatar, Typography, Popconfirm, Badge } from "antd";
import { DeleteOutlined } from "@ant-design/icons";
import dayjs from "dayjs";
import relativeTime from "dayjs/plugin/relativeTime";
import { Link } from "react-router-dom";
const { Paragraph } = Typography;
dayjs.extend(relativeTime);

export default function Conversations({ conversations }) {
	return (
		<List
			itemLayout="vertical"
			size="large"
			pagination={{
				onChange: (page) => {
					console.log(page);
				},
				pageSize: 5,
			}}
			dataSource={conversations}
			renderItem={(conversation) => (
				<List.Item
					key={conversation.id}
					actions={[
						<Popconfirm
							title="Sure to delete conversation?"
							onConfirm={() => console.log(conversation.id)}
						>
							<DeleteOutlined />
							<span className="converstion-action-delete"> Delete</span>
						</Popconfirm>,
					]}
				>
					<List.Item.Meta
						avatar={
							<Badge count={conversation.numNewMessage}>
								<Avatar src={conversation.partner.avatar} />
							</Badge>
						}
						title={
							<Link to={`${conversation.partner.id}`}>
								{conversation.partner.username}
							</Link>
						}
						description={dayjs.unix(conversation.lastMessage.created).fromNow()}
					/>
					<Paragraph strong={conversation.lastMessage.unread}>
						{conversation.lastMessage.message}
					</Paragraph>
				</List.Item>
			)}
		/>
	);
}
