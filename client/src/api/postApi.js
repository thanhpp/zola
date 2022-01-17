import axiosClient from "./axiosClient";

export async function getPostList() {
	const { data } = await axiosClient.get("/post");
	return data;
}

export async function addPost() {}
