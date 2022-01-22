import axiosClient from "./axiosClient";

export async function getPostList({ pageParam = 1 }) {
	const { data } = await axiosClient.get(`/admin/posts?index=${pageParam}`);
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

export async function getPostComment({ id, pageParam = 1 }) {
	const { data } = await axiosClient.get(
		`/post/${id}/comment?index=${pageParam}&count=10`
	);
	return {
		data,
		nextPage: pageParam + 1,
	};
}

export async function deletePostComment({ postId, commentId, comment }) {
	const commentData = new FormData();
	commentData.append("comment", comment);
	const { data } = await axiosClient.delete(
		`/post/${postId}/comment/${commentId}`,
		{
			data: commentData,
		}
	);
	return data;
}

export async function addPostComment({ id, comment }) {
	const { data } = await axiosClient.post(`/post/${id}/comment`, comment);
	return data;
}

export async function getPostMedia(url) {
	const response = await axiosClient.get(url, { responseType: "arraybuffer" });
	const base64Img = Buffer.from(response.data).toString("base64");
	return URL.createObjectURL(base64Img);
}

export async function deletePost(id) {
	const { data } = axiosClient.delete(`/post/${id}`);
	return data;
}

export async function likePost(id) {
	const { data } = await axiosClient.put(`/post/${id}/like`);
	return data;
}
