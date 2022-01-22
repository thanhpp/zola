import React, { useContext, useState } from "react";
import Conversations from "../../components/list/Conversations";
import { Button } from "antd";
import Spinner from "../../components/spinner/Spinner";
import ModalNewChat from "../../components/modal/ModalNewChat";
import { useQuery } from "react-query";
import { getConversationList } from "../../api/chatApi";
import AuthContext from "../../context/authContext";
import { useNavigate } from "react-router-dom";

export default function ConversationsList() {
	let navigate = useNavigate();
	const { user } = useContext(AuthContext);
	const [displayModal, setDisplayModal] = useState(false);
	const { data: conversations, isLoading } = useQuery("converstations", () =>
		getConversationList(user.userId)
	);

	if (isLoading) return <Spinner />;

	const onCreateChat = (values) => {
		navigate(`/messages/${values.userId}`);
	};

	return (
		<div>
			<Button type="primary" block onClick={() => setDisplayModal(true)}>
				New Message
			</Button>
			{conversations && <Conversations conversations={conversations.data} />}
			<ModalNewChat
				visible={displayModal}
				setVisible={setDisplayModal}
				onCreate={onCreateChat}
			/>
		</div>
	);
}
