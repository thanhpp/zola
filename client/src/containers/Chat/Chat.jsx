import React, { useState } from "react";
import Messages from "../../components/chat/Messages";
import Editor from "../../components/chat/Editor";
import Spinner from "../../components/spinner/Spinner";
import styles from "./Chat.module.css";
import ScrollToBottom from "react-scroll-to-bottom";
import dayjs from "dayjs";
import relativeTime from "dayjs/plugin/relativeTime";
import { Comment } from "antd";
dayjs.extend(relativeTime);

const conversation = {
	id: "2",
	messages: [
		{
			messageId: "1234345",
			message: "This is a messages",
			unread: 1, // đã đọc
			created: "1638785915",
			sender: {
				id: "41324124", // là user đang đăng nhập
				username: "John Doe",
				avatar: "https://joeschmoe.io/api/v1/random",
			},
		},
		{
			messageId: "1234335",
			message: "This is a messages",
			unread: 0, // chưa đọc
			created: "1638785965",
			sender: {
				id: "41324123",
				username: "John Doe",
				avatar: "https://joeschmoe.io/api/v1/random",
			},
		},
	],
	isBlocked: 0, // không bị block
};

export default function Chat() {
	const [messages, setMessages] = useState(conversation.messages);
	const [message, setMessage] = useState({
		submitting: false,
		value: "",
	});

	const [isLoading, setIsLoading] = useState(false);

	const handleChange = (e) => {
		setMessage({ ...message, value: e.target.value });
	};

	const handleSubmit = () => {
		if (!!message.submitting) {
			return;
		}

		setMessage({ ...message, submitting: true });

		setTimeout(() => {
			setMessage({
				...message,
				submitting: false,
				value: "",
			});

			setMessages([
				...messages,
				{
					author: "You",
					content: <p>{message.value}</p>,
					datetime: dayjs().fromNow(),
				},
			]);
		}, 1000);
	};

	return (
		<div className={styles.background}>
			{isLoading && <Spinner />}

			<ScrollToBottom className={styles.container}>
				{messages.length > 1 && <Messages messages={messages} />}
			</ScrollToBottom>

			<Comment
				content={
					<Editor
						onChange={handleChange}
						onSubmit={handleSubmit}
						submitting={message.submitting}
						value={message.value}
					/>
				}
			/>
		</div>
	);
}
