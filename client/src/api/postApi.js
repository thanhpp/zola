import axiosClient from "./axiosClient";

export async function getPostList({ pageParam = 1 }) {
	const { data } = await axiosClient.get(`/post?index=${pageParam}`);
	return {
		data,
		nextPage: pageParam + 1,
	};
}

export async function addPost(post) {
	const { data } = await axiosClient.post("/post", post);
	return data;
}

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

export async function getPostMedia(url) {
	const response = await axiosClient.get(url, { responseType: "arraybuffer" });
	const base64Img = Buffer.from(response.data).toString("base64");
	return base64Img;
}

export async function deletePost(id) {
	const { data } = axiosClient.delete(`/post/${id}`);
	return data;
}
