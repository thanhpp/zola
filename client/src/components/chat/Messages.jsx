import React, { useContext } from "react";
import "antd/dist/antd.css";
import { Avatar, Typography, Card } from "antd";
import { DeleteOutlined, UserOutlined } from "@ant-design/icons";
import dayjs from "dayjs";
import relativeTime from "dayjs/plugin/relativeTime";
import styles from "./Messages.module.css";
import AuthContext from "../../context/authContext";
dayjs.extend(relativeTime);
const { Meta } = Card;

export default function Messages(props) {
	const { user } = useContext(AuthContext);
	const { message, handleDelete } = props;

	return (
		<>
			{/* {messages.map((message) => { */}
			{/* return ( */}
			<Card
				key={message.message_id}
				className={
					message.sender.id === user.userId
						? styles["message-reciever-bubble"]
						: styles["message-sender-bubble"]
				}
				size="small"
			>
				<Meta
					avatar={
						message.sender.avatar ? (
							<Avatar src={message.sender.avatar} />
						) : (
							<Avatar icon={<UserOutlined />} />
						)
					}
					title={
						<div
							className={
								message.sender.id === user.userId
									? styles["message-reciever-name"]
									: styles["message-sender-name"]
							}
						>
							{message.sender.id === user.userId ? "You" : message.sender.name}
							<div
								className={styles["message-icon"]}
								onClick={() => handleDelete(message.message_id)}
							>
								<DeleteOutlined />
							</div>
						</div>
					}
					description={
						<p
							style={{
								color: message.sender.id === user.userId ? "white" : "black",
							}}
						>
							{dayjs.unix(message.created).fromNow()}
						</p>
					}
				/>
				<p></p>
				<Typography
					style={{
						color: message.sender.id === user.userId ? "white" : "black",
					}}
				>
					{message.message}
				</Typography>
			</Card>
			{/* ); */}
			{/* } */}
			{/* ) */}
			{/* } */}
		</>
	);
}
