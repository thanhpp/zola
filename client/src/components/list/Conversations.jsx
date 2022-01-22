import React from "react";
import "antd/dist/antd.css";
import { List, Avatar, Typography, Popconfirm } from "antd";
import { DeleteOutlined, UserOutlined } from "@ant-design/icons";
import dayjs from "dayjs";
import relativeTime from "dayjs/plugin/relativeTime";
import { Link } from "react-router-dom";
const { Paragraph } = Typography;
dayjs.extend(relativeTime);

export default function Conversations(props) {
	const { conversations, handleDelete } = props;
	return (
		<List
			itemLayout="vertical"
			size="large"
			pagination={{
				onChange: () => {
					//console.log(page);
				},
				pageSize: 20,
			}}
			dataSource={conversations}
			renderItem={(conversation) =>
				conversation.lastmessage.message ? (
					<List.Item
						key={conversation.id}
						actions={[
							<Popconfirm
								title="Sure to delete conversation?"
								onConfirm={() => handleDelete(conversation.id)}
							>
								<DeleteOutlined />
								<span className="converstion-action-delete"> Delete</span>
							</Popconfirm>,
						]}
					>
						<List.Item.Meta
							avatar={
								conversation.partner.avatar ? (
									<Avatar src={conversation.partner.avatar} />
								) : (
									<Avatar icon={<UserOutlined />} />
								)
							}
							title={
								<Link to={`${conversation.partner.id}`}>
									{conversation.partner.name}
								</Link>
							}
							description={dayjs
								.unix(conversation.lastmessage.created)
								.fromNow()}
						/>
						<Paragraph strong={+conversation.lastmessage.unread}>
							{conversation.lastmessage.message}
						</Paragraph>
					</List.Item>
				) : null
			}
		/>
	);
}
