import React, { useState } from "react";
import Post from "../components/user/Post";
import Comments from "../components/user/Comments";
import Editor from "../components/chat/Editor";
import { Comment, message } from "antd";
import { useQuery } from "react-query";
import { getPost, getPostComment } from "../api/postApi";
import { useParams } from "react-router-dom";
import Spinner from "../components/spinner/Spinner";

export default function PostDetail() {
	const [index, setIndex] = useState(1);
	const [enabled, setEnabled] = useState(true);
	const { id } = useParams();
	const { data: post, isLoading } = useQuery(["posts", id], () => getPost(id));
	const {
		data: comments,
		isLoading: isCommentsLoading,
		refetch,
	} = useQuery("comments", () => getPostComment({ id, index }), {
		enabled: enabled,
		retry: false,
		onSuccess: () => {
			setEnabled(false);
		},
		onError: (error) => {
			message.error({
				content: `Code: ${error.response?.data?.code};
				Message: ${error.response?.data?.message}`,
			});
		},
	});

	const onLoadMore = () => {
		setIndex(index + 1);
		refetch();
	};

	const [comment, setComment] = useState({
		submitting: false,
		value: "",
	});

	const handleChange = (e) => {
		setComment({ ...comment, value: e.target.value });
	};

	const handleSubmit = () => {
		if (!comment.value) return;
		setComment({ ...comment, submitting: true });
		console.log(comment.value);
		setTimeout(() => {
			setComment({
				submitting: false,
				value: "",
			});
		}, 1000);
	};

	if (isLoading) return <Spinner />;

	return (
		<>
			{post && (
				<Post post={post.data}>
					{comments && (
						<Comments
							comments={comments.data}
							isLoading={isCommentsLoading}
							onLoadMore={onLoadMore}
						/>
					)}
				</Post>
			)}
			{parseInt(post.data.can_comment) ? (
				<Comment
					content={
						<Editor
							onChange={handleChange}
							onSubmit={handleSubmit}
							submitting={comment.submitting}
							value={comment.value}
						/>
					}
				/>
			) : null}
		</>
	);
}
