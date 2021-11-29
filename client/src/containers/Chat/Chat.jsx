import React, { useState } from "react";
import Messages from "../../components/chat/Messages";
import Editor from "../../components/chat/Editor";
import Spinner from "../../components/spinner/Spinner";
import styles from "./Chat.module.css";
import ScrollToBottom from "react-scroll-to-bottom";
import dayjs from "dayjs";
import relativeTime from "dayjs/plugin/relativeTime";
dayjs.extend(relativeTime);

export default function Chat() {
	const [messages, setMessages] = useState([
		{
			author: "Han Solo",
			content: <p>this is a message</p>,
			datetime: dayjs().fromNow(),
		},
		{
			author: "Han Solo",
			content: <p>this is a message</p>,
			datetime: dayjs().fromNow(),
		},
		{
			author: "Han Solo",
			content: <p>this is a message</p>,
			datetime: dayjs().fromNow(),
		},
		{
			author: "Han Solo",
			content: <p>this is a message</p>,
			datetime: dayjs().fromNow(),
		},
		{
			author: "Han Solo",
			content: <p>this is a message</p>,
			datetime: dayjs().fromNow(),
		},
		{
			author: "Han Solo",
			content: <p>this is a message</p>,
			datetime: dayjs().fromNow(),
		},
		{
			author: "Han Solo",
			content: <p>this is a message</p>,
			datetime: dayjs().fromNow(),
		},
	]);
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
				<div className={styles.header}>Han Solo</div>
				{messages.length > 1 && <Messages messages={messages} />}
				<Editor
					onChange={handleChange}
					onSubmit={handleSubmit}
					submitting={message.submitting}
					value={message.value}
				/>
			</ScrollToBottom>
		</div>
	);
}
