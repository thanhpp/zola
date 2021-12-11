import React, { useState } from "react";
import Conversations from "../../components/list/Conversations";
import { Button } from "antd";
import ModalNewChat from "../../components/modal/ModalNewChat";

const conversations = [
	{
		id: "2",
		partner: {
			id: "123123",
			username: "John Doe",
			avatar: "https://joeschmoe.io/api/v1/random",
		},
		lastMessage: {
			message: "this is a message",
			created: "1638785955",
			unread: 0, // 0 : read; 1: unread
		},
		numNewMessage: "2",
	},
	{
		id: "4",
		partner: {
			id: "123123",
			username: "John Doe",
			avatar: "https://joeschmoe.io/api/v1/random",
		},
		lastMessage: {
			message: "this is a message",
			created: "1638785915",
			unread: 1,
		},
		numNewMessage: "0",
	},
];

export default function ConversationsList() {
	const [displayModal, setDisplayModal] = useState(false);
	return (
		<div>
			<Button type="primary" block onClick={() => setDisplayModal(true)}>
				New Message
			</Button>
			<Conversations conversations={conversations} />
			<ModalNewChat visible={displayModal} setVisible={setDisplayModal} />
		</div>
	);
}
