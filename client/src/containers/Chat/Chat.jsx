import React, { useContext, useEffect, useState } from "react";
import Messages from "../../components/chat/Messages";
import Editor from "../../components/chat/Editor";
//import Spinner from "../../components/spinner/Spinner";
import styles from "./Chat.module.css";
import ScrollToBottom from "react-scroll-to-bottom";
import dayjs from "dayjs";
import relativeTime from "dayjs/plugin/relativeTime";
import { Comment, message } from "antd";
import io from "socket.io-client";
import { useParams } from "react-router-dom";
import AuthContext from "../../context/authContext";
dayjs.extend(relativeTime);

const socket = io.connect(process.env.REACT_APP_CHAT_URL, {
	reconnection: false,
});

export default function Chat(props) {
	//const { receiverId } = props;
	const { user } = useContext(AuthContext);
	const { id } = useParams();
	const [isLoading, setIsLoading] = useState(false);
	const [messages, setMessages] = useState([]);
	const [chatMessage, setChatMessage] = useState({
		message_id: "",
		event: "joinchat",
		sender: user.userId,
		receiver: 41324124,
		created: Date.now(),
		content: "",
	});

	const handleChange = (e) => {
		setChatMessage({
			...message,
			...{
				message_id: "",
				event: "send",
				sender: user.userId,
				receiver: 41324124,
				created: Date.now(),
				content: e.target.value,
			},
		});
	};

	useEffect(() => {
		socket.emit("joinchat", chatMessage, (error) => {
			message.error(error);
		});

		socket.on("connection_timeout", (err) => {
			console.log(err);
		});

		socket.on("connect_error", (err) => {
			console.log(err);
		});

		socket.on("reconnecting", () => {
			console.log("trying to recconect");
		});

		socket.on("reconnect_attempt", () => {
			console.log("trying to recconect with many attempts ...");
		});

		socket.on("disconnect", (reason) => {
			console.log("socket connection disconnected", reason);
		});

		socket.on("onmessage", (incommingMessage) => {
			setMessages((messages) => [...messages, incommingMessage]);
		});
		return () => {
			socket.disconnect();
		};
	}, [chatMessage]);

	const handleSubmit = () => {
		if (!chatMessage.content) {
			return;
		}
		setIsLoading(true);
		socket.emit("send", chatMessage);
		setIsLoading(false);
	};

	return (
		<div className={styles.background}>
			<ScrollToBottom className={styles.container}>
				{messages.length > 1 && <Messages messages={messages} />}
			</ScrollToBottom>

			<Comment
				content={
					<Editor
						onChange={handleChange}
						onSubmit={handleSubmit}
						submitting={isLoading}
						value={chatMessage}
					/>
				}
			/>
		</div>
	);
}
