import axiosClient from "./axiosClient";

export async function getPostList() {
	const { data } = await axiosClient.get("/post");
	return data;
}

export async function addPost() {}

export async function getPost(id) {
	const { data } = await axiosClient.get(`/post/${id}`);
	return data;
}

export async function getPostComment({ id, index }) {
	const { data } = await axiosClient.get(
		`/post/${id}/comment?index=${index}&count=10`
	);
	return data;
}
