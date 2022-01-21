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
//import { socket } from "../../api/socket";
import AuthContext from "../../context/authContext";
dayjs.extend(relativeTime);

let socket = new WebSocket(process.env.REACT_APP_CHAT_URL);
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
		receiver: id,
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
				receiver: id,
				created: Date.now(),
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
			console.log(e);
			console.log(e.data);
			console.log(JSON.parse(e.data));
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

		console.log(socket.readyState);

		//close when unmount
		return () => {
			socket.close();
		};
	}, []);

	const handleSubmit = () => {
		if (!chatMessage.content) {
			return;
		}
		setIsLoading(true);
		//send message
		if (socket.readyState !== 1) {
			message.error("error when connect with websocket");
		}
		socket.send(JSON.stringify(chatMessage));
		setIsLoading(false);
		setChatMessage({
			...message,
			...{
				message_id: "",
				event: "send",
				sender: user.userId,
				receiver: "0fc9ef71-708e-11ec-bd01-0242c0a83003",
				created: Date.now(),
				content: "",
			},
		});
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
