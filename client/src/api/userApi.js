import axiosClient from "./axiosClient";

export async function getUserList() {
	const { data } = await axiosClient.get("/admin/users");
	return data;
}

export async function setUserState() {}

export async function deleteUser(id) {
	console.log(id);
	//const {data} = await axiosClient.delete()
}

export async function getUserInfo(id) {
	const res = axiosClient.get(`/user/${id}`);
	return res;
}
