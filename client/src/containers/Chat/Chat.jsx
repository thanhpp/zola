import React, { useContext, useEffect, useState } from "react";
import Messages from "../../components/chat/Messages";
import Editor from "../../components/chat/Editor";
//import Spinner from "../../components/spinner/Spinner";
import styles from "./Chat.module.css";
import ScrollToBottom from "react-scroll-to-bottom";
import dayjs from "dayjs";
import relativeTime from "dayjs/plugin/relativeTime";
import { Comment, message } from "antd";
import { useParams } from "react-router-dom";
import { socket } from "../../api/socket";
import AuthContext from "../../context/authContext";
dayjs.extend(relativeTime);

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
		//connect to socket
		socket.onopen = () => {
			console.log("connected, bitches");
		};

		//receive message
		socket.onmessage = (e) => {
			console.log(e);
			console.log(e.data);
		};

		//error
		socket.onerror = (error) => {
			console.log(error);
		};
		return () => {
			socket.close();
		};
	}, [chatMessage]);

	const handleSubmit = () => {
		if (!chatMessage.content) {
			return;
		}
		setIsLoading(true);
		//send message
		socket.send(chatMessage);
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
						value={chatMessage.content}
					/>
				}
			/>
		</div>
	);
}
