import React, { useContext, useState } from "react";
import {
	useInfiniteQuery,
	useMutation,
	useQueryClient,
	useQuery,
} from "react-query";
import AuthContext from "../../context/authContext";
import { useParams } from "react-router-dom";
import Chat from "./Chat";
import { getConversation, deleteMessage } from "../../api/chatApi";
import { getUserChatInfo } from "../../api/userApi";
import { message } from "antd";
import Spinner from "../../components/spinner/Spinner";

export default function WrapperChat() {
	const queryClient = useQueryClient();
	const [enabled, setEnabled] = useState(true);
	const [chatHistory, setChatHistory] = useState([]);
	const { id } = useParams();
	const { user } = useContext(AuthContext);
	const { userId } = user;
	const socket = new WebSocket(
		`${process.env.REACT_APP_CHAT_URL}?id=${userId}`
	);

	const {
		data: messages,
		fetchNextPage,
		hasNextPage,
		isLoading,
	} = useInfiniteQuery(
		"messages",
		({ pageParam = 1 }) => getConversation({ id, pageParam }),
		{
			enabled: enabled,
			getNextPageParam: (lastPage) => {
				if (lastPage.data.data.conversation.length !== 0)
					return lastPage.nextPage;
				return undefined;
			},
			onError: (error) => {
				message.error({
					content: `Code: ${error.response.data.code};
				Message: ${error.response.data.message}`,
				});
			},
			onSuccess: (data) => {
				setEnabled(false);
				//console.log(data.pages[data.pages.length - 1].data.data.conversation);
				setChatHistory((chat) => [
					...chat,
					...data.pages[data.pages.length - 1].data.data.conversation,
				]);
			},
		}
	);

	const { data: currentUser } = useQuery(
		["messages", "currentUser", user.userId],
		() => getUserChatInfo(user.userId),
		{
			enabled: enabled,
		}
	);
	const { data: partnerInfo } = useQuery(
		["messages", "partner", id],
		() => getUserChatInfo(id),
		{
			enabled: enabled,
		}
	);

	const { mutate: handleDelete } = useMutation(deleteMessage, {
		onSuccess: () => {
			//console.log(data);
			queryClient.invalidateQueries("messages");
		},
		onError: (error) => {
			console.log(error);
			message.error({
				content: `Code: ${error.response.data.code};
				Message: ${error.response.data.message}`,
			});
		},
		onMutate: () => {
			message.loading("deleting...");
		},
	});

	const transformMessage = (message) => {
		const newMessage = {
			sender: {},
		};
		const { content, sender, message_id, created } = message;
		newMessage.message_id = message_id;
		newMessage.message = content;
		newMessage.created = created;
		newMessage.sender.id = sender;
		if (sender === user.userId) {
			newMessage.sender.avatar = currentUser?.avatar || "";
			newMessage.sender.username = currentUser.username;
			newMessage.sender.name = currentUser.name;
		} else {
			newMessage.sender.avatar = partnerInfo.avatar;
			newMessage.sender.username = partnerInfo.username;
			newMessage.sender.name = partnerInfo.name;
		}

		setChatHistory((chat) => [newMessage, ...chat]);
	};

	if (isLoading) return <Spinner />;
	//console.log(chatHistory);

	return (
		<>
			{messages && (
				<Chat
					socket={socket}
					user={user}
					id={id}
					fetchNextPage={fetchNextPage}
					hasNextPage={hasNextPage}
					chat={chatHistory}
					onCreate={transformMessage}
					handleDelete={handleDelete}
				/>
			)}
		</>
	);
}
