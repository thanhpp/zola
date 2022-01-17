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

export async function getUserFriend({ pageParam = 1 }) {
	const { data } = await axiosClient.get(`/friend?index=${pageParam}`);
	return {
		data,
		nextPage: pageParam + 1,
	};
}
