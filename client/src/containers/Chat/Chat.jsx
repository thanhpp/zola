import React, { useEffect, useState } from "react";
import Messages from "../../components/chat/Messages";
import Editor from "../../components/chat/Editor";
import styles from "./Chat.module.css";
import InfiniteScroll from "react-infinite-scroll-component";
import dayjs from "dayjs";
import { Comment, message } from "antd";

export default function Chat(props) {
	const {
		socket,
		user,
		id,
		chat,
		onCreate,
		hasNextPage,
		fetchNextPage,
		handleDelete,
	} = props;
	const [isLoading, setIsLoading] = useState(false);
	//const [conversation, setConversation] = useState([]);
	const [chatMessage, setChatMessage] = useState({
		message_id: "",
		event: "joinchat",
		sender: user.userId,
		receiver: id,
		created: `${dayjs().unix()}`,
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
				created: `${dayjs().unix()}`,
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
			console.log(JSON.parse(e.data));
			onCreate(JSON.parse(e.data));
			//return false;
			//setConversation((messages) => [...messages, JSON.parse(e.data)]);
		};

		//error
		socket.onerror = (error) => {
			message.error("Error in connection with Websocket");
			console.log(error);
		};

		//close
		socket.onclose = () => {
			console.log("connection closed");
		};

		if (socket.readyState === 0) {
			message.loading("connecting to websocket");
		}

		//console.log(chat);

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
					created: `${dayjs().unix()}`,
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
			<div id="scrollableDiv" className={styles.container}>
				<InfiniteScroll
					dataLength={chat.length}
					//next={fetchNextPage}
					next={() => console.log("called next")}
					style={{ display: "flex", flexDirection: "column-reverse" }}
					inverse={true}
					hasMore={hasNextPage}
					loader={"loading..."}
					initialScrollY={800}
					scrollableTarget="scrollableDiv"
				>
					{chat.map((message) => {
						return (
							<Messages
								key={message.message_id}
								message={message}
								handleDelete={handleDelete}
							/>
						);
					})}
					{/* <Messages messages={chat} /> */}
				</InfiniteScroll>
			</div>

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
