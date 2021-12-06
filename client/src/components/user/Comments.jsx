import React from "react";
import "antd/dist/antd.css";
import { Comment, Avatar, Tooltip, Typography } from "antd";
import { DeleteOutlined } from "@ant-design/icons";

export default function Comments({ comments }) {
	const actions = [
		<Tooltip key="comment-basic-delete" title="Delete comment">
			<DeleteOutlined />
			<span> Delete comment</span>
		</Tooltip>,
	];
	return (
		<>
			{comments.map((comment) => {
				return (
					<Comment
						key={comment.id}
						actions={actions}
						author={comment.author}
						avatar={<Avatar src={comment.avatar} alt="avatar" />}
						content={
							<Typography.Paragraph>{comment.content}</Typography.Paragraph>
						}
					/>
				);
			})}
		</>
	);
}
