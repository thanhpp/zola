import React from "react";
import "antd/dist/antd.css";
import { Avatar, Typography, Card } from "antd";
import { DeleteOutlined } from "@ant-design/icons";
import dayjs from "dayjs";
import relativeTime from "dayjs/plugin/relativeTime";
import styles from "./Messages.module.css";
dayjs.extend(relativeTime);
const { Meta } = Card;

export default function Messages({ messages }) {
	//const userId = "41324124";
	console.log(messages);
	return (
		<>
			{/* {messages.map((message) => {
				return (
					<Card
						key={message.messageId}
						className={
							message.sender.id === userId
								? styles["message-reciever-bubble"]
								: styles["message-sender-bubble"]
						}
						size="small"
					>
						<Meta
							avatar={<Avatar src={message.sender.avatar} />}
							title={
								<div
									className={
										message.sender.id === userId
											? styles["message-reciever-name"]
											: styles["message-sender-name"]
									}
								>
									{message.sender.id === userId
										? "You"
										: message.sender.username}
									<div
										className={styles["message-icon"]}
										onClick={() => console.log(message.messageId)}
									>
										<DeleteOutlined />
									</div>
								</div>
							}
							description={
								<p
									style={{
										color: message.sender.id === userId ? "white" : "black",
									}}
								>
									{dayjs.unix(message.created).fromNow()}
								</p>
							}
						/>
						<p></p>
						<Typography
							style={{
								color: message.sender.id === userId ? "white" : "black",
							}}
						>
							{message.message}
						</Typography>
					</Card>
				);
			})} */}
			{/* {messages.map((message) => {
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
			})} */}
		</>
	);
}
