import axios from "axios";
const url = "https://zola.thanhpp.ninja";

export async function loginUser(user) {
	const { phonenumber, password } = user;
	const res = await axios.post(`${url}/login`, {
		phonenumber,
		password,
	});
	return res.data;
}

export async function logoutUser() {
	return await axios.get(`${url}/logout`, {
		headers: {
			Authorization: `Bearer ${localStorage.getItem("token")}`,
		},
	});
}

export async function signUpUser(user) {
	const { phonenumber, password } = user;
	const res = await axios.post(`${process.env.REACT_APP_API_URL}/signup`, {
		phonenumber,
		password,
	});
	return res.data;
}
