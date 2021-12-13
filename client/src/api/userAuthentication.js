import axios from "axios";
const baseUrl = "http://127.0.0.1:12100";

export async function loginUser(phoneNumber, password) {
	const formData = new FormData();
	formData.append("phonenumber", phoneNumber);
	formData.append("password", password);
	const res = await axios.post(`${baseUrl}/login`, formData);
	return res.data;
}

export async function logoutUser() {
	return await axios.get(`${baseUrl}/logout`);
}
