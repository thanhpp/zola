import React, { useState } from "react";
import Post from "../components/user/Post";
import Comments from "../components/user/Comments";
import Editor from "../components/chat/Editor";
import { Comment, Avatar } from "antd";

const post = {
	author: "John Doe",
	avatar: "https://joeschmoe.io/api/v1/random",
	media: [
		// "https://gw.alipayobjects.com/zos/rmsportal/mqaQswcyDLcXyDKnZfES.png",
		// "https://gw.alipayobjects.com/zos/rmsportal/mqaQswcyDLcXyDKnZfES.png",
		// "https://gw.alipayobjects.com/zos/rmsportal/mqaQswcyDLcXyDKnZfES.png",
		// "https://gw.alipayobjects.com/zos/rmsportal/mqaQswcyDLcXyDKnZfES.png",
		"https://archive.org/download/BigBuckBunny_124/Content/big_buck_bunny_720p_surround.mp4",
	],
	content:
		"We supply a series of design principles, practical patterns and high quality design resources (Sketch and Axure),\
        to help people create their product prototypes beautifully and efficiently.We supply a series of design principles, practical patterns and high quality design resources (Sketch and Axure),\
        to help people create their product prototypes beautifully and efficiently.We supply a series of design principles, practical patterns and high quality design resources (Sketch and Axure),\
        to help people create their product prototypes beautifully and efficiently.We supply a series of design principles, practical patterns and high quality design resources (Sketch and Axure),\
        to help people create their product prototypes beautifully and efficiently.We supply a series of design principles, practical patterns and high quality design resources (Sketch and Axure),\
        to help people create their product prototypes beautifully and efficiently.",
	like: 126,
	comment: 126,
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
		</>
	);
}
