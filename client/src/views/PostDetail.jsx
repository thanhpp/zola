import React, { useState } from "react";
import Post from "../components/user/Post";
import Comments from "../components/user/Comments";
import Editor from "../components/chat/Editor";
import { Comment, message } from "antd";
import {
	useQuery,
	useMutation,
	useQueryClient,
	useInfiniteQuery,
} from "react-query";
import {
	getPost,
	getPostComment,
	addPostComment,
	deletePostComment,
	likePost,
} from "../api/postApi";
import { useParams } from "react-router-dom";
import Spinner from "../components/spinner/Spinner";

export default function PostDetail() {
	const queryClient = useQueryClient();
	const { id } = useParams();
	const [comment, setComment] = useState({
		value: "",
	});
	const { data: post, isLoading } = useQuery(["posts", id], () => getPost(id));
	const {
		data: comments,
		isLoading: isCommentsLoading,
		hasNextPage,
		fetchNextPage,
	} = useInfiniteQuery(
		["posts", id, "comments"],
		({ pageParam = 1 }) => getPostComment({ id, pageParam }),
		{
			//enabled: enabled,
			retry: false,
			getNextPageParam: (lastPage) => {
				//console.log(lastPage);
				if (lastPage.data.data.length !== 0) return lastPage.nextPage;
				return undefined;
			},
			onError: (error) => {
				message.error({
					content: `Code: ${error.response?.data?.code};
				Message: ${error.response?.data?.message}`,
				});
			},
		}
	);

	const { mutate: addNewComment, isLoading: isCommentSubmitting } = useMutation(
		addPostComment,
		{
			onSuccess: () => {
				//queryClient.invalidateQueries(["posts", id, "comments"]);
				setComment({ value: "" });
				queryClient.refetchQueries(["posts", id, "comments"]);
			},
			onError: (error) => {
				message.error({
					content: `Code: ${error.response?.data?.code};
				Message: ${error.response?.data?.message}`,
				});
			},
		}
	);

	const { mutate: deleteOldComment } = useMutation(deletePostComment, {
		onSuccess: () => {
			queryClient.invalidateQueries(["posts", id, "comments"]);
		},
		onError: (error) => {
			message.error({
				content: `Code: ${error.response?.data?.code};
				Message: ${error.response?.data?.message}`,
			});
		},
		onMutate: () => {
			message.loading("loading");
		},
	});

	const { mutate: postInteraction } = useMutation(likePost, {
		onSuccess: () => {
			queryClient.invalidateQueries(["posts", id]);
		},
		onError: (error) => {
			message.error({
				content: `Code: ${error.response?.data?.code};
				Message: ${error.response?.data?.message}`,
			});
		},
		onMutate: () => {
			message.loading("loading");
		},
	});

	const handleChange = (e) => {
		setComment({ ...comment, value: e.target.value });
	};

	const handleSubmit = () => {
		if (!comment.value) return;
		const formData = new FormData();
		formData.append("comment", comment.value);
		addNewComment({ id: id, comment: formData });
	};

	if (isLoading) return <Spinner />;

	return (
		<>
			{post && (
				<Post post={post.data} handleInteraction={postInteraction}>
					{comments && (
						<Comments
							comments={comments}
							isLoading={isCommentsLoading}
							onLoadMore={fetchNextPage}
							handleDeleteComment={deleteOldComment}
							postId={id}
							hasMoreComment={hasNextPage}
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
							submitting={isCommentSubmitting}
							value={comment.value}
						/>
					}
				/>
			) : null}
		</>
	);
}
