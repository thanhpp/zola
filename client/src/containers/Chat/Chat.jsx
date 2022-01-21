import React, { useContext, useEffect, useState } from "react";
import Messages from "../../components/chat/Messages";
import Editor from "../../components/chat/Editor";
import styles from "./Chat.module.css";
import ScrollToBottom from "react-scroll-to-bottom";
import dayjs from "dayjs";
import relativeTime from "dayjs/plugin/relativeTime";
import { Comment, message } from "antd";
import { useParams } from "react-router-dom";
//import { socket } from "../../api/socket";
import AuthContext from "../../context/authContext";
dayjs.extend(relativeTime);

function Chat(props) {
	const { socket, user } = props;
	const { id } = useParams();
	const [isLoading, setIsLoading] = useState(false);
	const [messages, setMessages] = useState([]);
	const [chatMessage, setChatMessage] = useState({
		message_id: "",
		event: "joinchat",
		sender: user.userId,
		receiver: id,
		created: Date.now().toString(),
		content: "",
	});

	const handleChange = (e) => {
		setChatMessage({
			...message,
			...{
				message_id: "",
				event: "send",
				sender: user.userId,
				receiver: id,
				created: Date.now().toString(),
				content: e.target.value,
			},
		});
	};

	useEffect(() => {
		//connect to socket
		socket.onopen = () => {
			console.log("connected to websocket");
			socket.send(JSON.stringify(chatMessage));
		};

		//receive message
		socket.onmessage = (e) => {
			//console.log(e);
			//console.log(e.data);
			console.log(JSON.parse(e.data));
			return false;
			//setMessages((messages) => [...messages, e.data]);
		};

		//error
		socket.onerror = (error) => {
			console.log(error);
		};

		//close

		socket.onclose = () => {
			console.log("connection closed");
			// socket = new WebSocket(process.env.REACT_APP_CHAT_URL);
			// console.log("connection openned");
		};

		//console.log(socket.readyState);

		if (socket.readyState === 0) {
			message.loading("connecting to websocket");
		}
		//close when unmount
		return () => {
			socket.close();
		};
	}, []);

	const handleSubmit = () => {
		try {
			if (!chatMessage.content) {
				return;
			}
			setIsLoading(true);
			//send message
			socket.send(JSON.stringify(chatMessage));
			setIsLoading(false);
			setChatMessage({
				...message,
				...{
					message_id: "",
					event: "send",
					sender: user.userId,
					receiver: id,
					created: Date.now().toString(),
					content: "",
				},
			});
		} catch (error) {
			console.log(error);
			setIsLoading(false);
			message.error("error when connect with websocket");
		}
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

export default function WrapperChat() {
	const { user } = useContext(AuthContext);
	const { userId } = user;
	const socket = new WebSocket(
		`${process.env.REACT_APP_CHAT_URL}?id=${userId}`
	);
	return <Chat socket={socket} user={user} />;
}
