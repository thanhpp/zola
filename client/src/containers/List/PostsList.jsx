import React, { useState } from "react";
import Posts from "../../components/list/Posts";
import { Button, message } from "antd";
import ModalNewPost from "../../components/modal/ModalFormPost";
import { useMutation, useInfiniteQuery, useQueryClient } from "react-query";
import { getPostList, addPost, deletePost } from "../../api/postApi";
import Spinner from "../../components/spinner/Spinner";

export default function PostsList() {
	const queryClient = useQueryClient();
	const [displayModal, setDisplayModal] = useState(false);
	const { data, isLoading, fetchNextPage, hasNextPage } = useInfiniteQuery(
		"posts",
		getPostList,
		{
			getNextPageParam: (lastPage) => {
				if (lastPage.data.data.posts.length > 20) return lastPage.nextPage;
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

	const { mutate: addNewPost } = useMutation(addPost, {
		onSuccess: () => {
			queryClient.invalidateQueries("posts");
			//queryClient.refetchQueries("posts");
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

	const { mutate: deleteOldPost } = useMutation(deletePost, {
		onSuccess: () => {
			queryClient.invalidateQueries("posts");
			message.success("Post deleted!");
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

	const onCreate = (values) => {
		addNewPost(values);
	};

	if (isLoading) return <Spinner />;

	return (
		<>
			<Button type="primary" block onClick={() => setDisplayModal(true)}>
				Add new post
			</Button>
			{data.pages && (
				<Posts
					pages={data.pages}
					fetchNextPage={fetchNextPage}
					hasNextPage={hasNextPage}
					handleDelete={deleteOldPost}
				/>
			)}
			<ModalNewPost
				visible={displayModal}
				onCreate={onCreate}
				setVisible={setDisplayModal}
			/>
		</>
	);
}
