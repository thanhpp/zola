import axiosClient from "./axiosClient";

export async function getUserList() {
	const { data } = await axiosClient.get("/admin/users");
	return data;
}

export async function setUserState(values) {
	const { user_id, state } = values;
	const formData = new FormData();
	formData.append("state", state);
	const { data } = await axiosClient.put(
		`/admin/users/${user_id}/state`,
		formData
	);
	return data;
}

export async function deleteUser(id) {
	const { data } = await axiosClient.delete(`/admin/users/${id}`);
	return data;
}

export async function getUserInfo(id) {
	const res = await axiosClient.get(`/user/${id}`);
	return res;
}

export async function editUserInfo(formData) {
	const { data } = await axiosClient.put(`/user`, formData);
	return data;
}

export async function getMyFriend({ pageParam = 1 }) {
	const { data } = await axiosClient.get(
		`/friend?index=${pageParam}&count=100`
	);
	return {
		data,
		nextPage: pageParam + 1,
	};
}

export async function getUserFriend(id) {
	const { data } = await axiosClient.get(`/friend/${id}`);
	return {
		data,
	};
}

export async function getUserChatInfo(id) {
	const { data } = await axiosClient.get(`/internal/user/${id}`);
	return data;
}
