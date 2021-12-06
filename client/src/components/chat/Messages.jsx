import React from "react";
import "antd/dist/antd.css";
import { Comment, Avatar, Typography } from "antd";
import dayjs from "dayjs";
import relativeTime from "dayjs/plugin/relativeTime";
dayjs.extend(relativeTime);

export default function Messages({ messages }) {
	const userId = "41324124";
	return (
		<>
			{messages.map((message) => {
				return (
					<Comment
						key={message.messageId}
						actions={[
							<span onClick={() => console.log(message.messageId)}>
								Delete
							</span>,
						]}
						author={
							message.sender.id === userId ? message.sender.username : "You"
						}
						avatar={<Avatar src={message.sender.avatar} alt="avatar" />}
						content={
							<Typography.Paragraph>{message.message}</Typography.Paragraph>
						}
						datetime={dayjs.unix(message.created).fromNow()}
					/>
				);
			})}
		</>
	);
}
