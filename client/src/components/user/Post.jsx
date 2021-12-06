import React from "react";
import "antd/dist/antd.css";
import { Comment, Tooltip, Avatar, Image, Space, Typography } from "antd";
import { LikeOutlined, MessageOutlined, LikeFilled } from "@ant-design/icons";

export default function Post(props) {
	const { post } = props;
	const actions = [
		<Tooltip key="comment-basic-like" title="Like">
			<span>
				{React.createElement(LikeOutlined)}
				<span className="comment-action">{post.like}</span>
			</span>
		</Tooltip>,
		<Tooltip key="comment-basic" title="Comment">
			<span>
				{React.createElement(MessageOutlined)}
				<span className="comment-action">{post.comment}</span>
			</span>
		</Tooltip>,
	];

	return (
		<Comment
			actions={actions}
			author={post.author}
			avatar={
				<Avatar src="https://joeschmoe.io/api/v1/random" alt="Han Solo" />
			}
			content={
				<>
					<Typography.Paragraph>{post.content}</Typography.Paragraph>
					{post.media[0]?.includes(".jpg") ||
					post.media[0]?.includes(".png") ? (
						<Image.PreviewGroup>
							<Space size={"large"}>
								{post.media.map((image) => {
									return <Image key={image} width={300} src={image} />;
								})}
							</Space>
						</Image.PreviewGroup>
					) : (
						<video width="600" controls>
							<source src={post.media[0]} type="video/mp4" />
						</video>
					)}
				</>
			}
			// datetime={
			// 	<Tooltip title={moment().format("YYYY-MM-DD HH:mm:ss")}>
			// 		<span>{moment().fromNow()}</span>
			// 	</Tooltip>
			// }
		>
			{props.children}
		</Comment>
	);
}
