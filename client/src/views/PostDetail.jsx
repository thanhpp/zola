import React, { useState } from "react";
import Post from "../components/user/Post";
import Comments from "../components/user/Comments";
import Editor from "../components/chat/Editor";
import { Comment, Avatar } from "antd";

const post = {
	id: "0",
	created: "1639036932",
	modified: "",
	author: {
		id: "123443245",
		name: "John Doe",
		avatar: "https://joeschmoe.io/api/v1/random",
	},
	// video: {
	// 	thumb:
	// 		"https://gw.alipayobjects.com/zos/rmsportal/mqaQswcyDLcXyDKnZfES.png",
	// 	url: "https://archive.org/download/BigBuckBunny_124/Content/big_buck_bunny_720p_surround.mp4",
	// },
	image: [
		{
			id: "2",
			url: "https://gw.alipayobjects.com/zos/rmsportal/mqaQswcyDLcXyDKnZfES.png",
		},
		{
			id: "3",
			url: "https://gw.alipayobjects.com/zos/rmsportal/mqaQswcyDLcXyDKnZfES.png",
		},
		{
			id: "4",
			url: "https://gw.alipayobjects.com/zos/rmsportal/mqaQswcyDLcXyDKnZfES.png",
		},
		{
			id: "5",
			url: "https://gw.alipayobjects.com/zos/rmsportal/mqaQswcyDLcXyDKnZfES.png",
		},
	],
	described: "This is a post",
	like: "126",
	comment: "126",
	is_liked: "1", // 0: not liked, 1: is liked
	is_blocked: 0, //0: not blocked, 1: is blocked
	can_edit: 1, //0: can not edit, 1: can edit
	can_comment: "0", //0: can't, 1: can
};

const comments = [
	{
		id: "2",
		author: "Jane Doe",
		avatar: "https://joeschmoe.io/api/v1/random",
		content: "horrah",
	},
	{
		id: "3",
		author: "Ann Doe",
		avatar: "https://joeschmoe.io/api/v1/random",
		content: "horrah",
	},
];

export default function PostDetail() {
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

	return (
		<>
			<Post post={post}>
				<Comments comments={comments} />
			</Post>
			{parseInt(post.can_comment) ? (
				<Comment
					avatar={
						<Avatar src="https://joeschmoe.io/api/v1/random" alt="Han Solo" />
					}
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
