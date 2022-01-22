import axios from "axios";
import axiosClient from "./axiosClient";

export async function loginUser(user) {
	const { phonenumber, password } = user;
	const res = await axios.post(`${process.env.REACT_APP_API_URL}/login`, {
		phonenumber,
		password,
	});
	return res.data;
}

export async function logoutUser() {
	return await axiosClient.get("/logout");
}

export async function signUpUser(user) {
	const { phonenumber, password } = user;
	const res = await axios.post(`${process.env.REACT_APP_API_URL}/signup`, {
		phonenumber,
		password,
	});
	return res.data;
}
