import React from "react";
import Conversations from "../../components/list/Conversations";
import { Button } from "antd";

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
			unread: 0,
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
	return (
		<div>
			<Button type="primary" block>
				New Message
			</Button>
			<Conversations conversations={conversations} />
		</div>
	);
}
