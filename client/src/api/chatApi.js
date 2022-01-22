import axiosChatClient from "./axiosChatClient";

export async function getConversationList(id) {
	const { data } = await axiosChatClient.get(
		`${process.env.REACT_APP_CHAT_ENDPOINT}/conversations?requestor_id=${id}`
	);
	return data;
}

export async function deleteConversation(id) {
	const { data } = await axiosChatClient.delete(
		`${process.env.REACT_APP_CHAT_ENDPOINT}/conversations/${id}`
	);
	return data;
}
