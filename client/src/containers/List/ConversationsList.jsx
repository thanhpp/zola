import React, { useContext, useState } from "react";
import Conversations from "../../components/list/Conversations";
import { Button, message } from "antd";
import Spinner from "../../components/spinner/Spinner";
import ModalNewChat from "../../components/modal/ModalNewChat";
import {
	useQuery,
	useInfiniteQuery,
	useMutation,
	useQueryClient,
} from "react-query";
import { getConversationList } from "../../api/chatApi";
import AuthContext from "../../context/authContext";
import { useNavigate } from "react-router-dom";
import { getMyFriend } from "../../api/userApi";
import { deleteConversation } from "../../api/chatApi";

export default function ConversationsList() {
	let navigate = useNavigate();
	const queryClient = useQueryClient();
	const { user } = useContext(AuthContext);
	const [displayModal, setDisplayModal] = useState(false);
	const { data: conversations, isLoading } = useQuery("converstations", () =>
		getConversationList(user.userId)
	);

	const { data: friends, isLoading: isFetchingFriend } = useInfiniteQuery(
		"myFriends",
		getMyFriend,
		{
			getNextPageParam: (lastPage) => {
				if (lastPage.data.data.friends !== null) return lastPage.nextPage;
				return undefined;
			},
			onError: (error) => {
				message.error({
					content: `Code: ${error.response.data.code};
				Message: ${error.response.data.message}`,
				});
			},
		}
	);

	const { mutate: deleteChat } = useMutation(deleteConversation, {
		onSuccess: () => {
			queryClient.invalidateQueries("converstations");
		},
		onError: (error) => {
			message.error({
				content: `Code: ${error.response.data.code};
				Message: ${error.response.data.message}`,
			});
		},
		onMutate: () => {
			message.loading("deleting conversation");
		},
	});

	if (isLoading) return <Spinner />;

	const onCreateChat = (values) => {
		navigate(`/messages/${values.userId}`);
	};

	return (
		<div>
			<Button type="primary" block onClick={() => setDisplayModal(true)}>
				New Message
			</Button>
			{conversations && (
				<Conversations
					conversations={conversations.data}
					handleDelete={deleteChat}
				/>
			)}
			{friends && (
				<ModalNewChat
					visible={displayModal}
					setVisible={setDisplayModal}
					onCreate={onCreateChat}
					friends={friends.pages[0].data.data.friends}
					isLoading={isFetchingFriend}
				/>
			)}
		</div>
	);
}
