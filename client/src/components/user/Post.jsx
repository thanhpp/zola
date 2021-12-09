import React from "react";
import "antd/dist/antd.css";
import { Comment, Tooltip, Avatar, Image, Space, Typography } from "antd";
import { LikeOutlined, MessageOutlined, LikeFilled } from "@ant-design/icons";

export default function Post(props) {
	const { post } = props;
	const { is_liked, like, comment, author, described } = post;
	const actions = [
		<Tooltip key="comment-basic-like" title="Like">
			<span>
				{React.createElement(parseInt(is_liked) ? LikeFilled : LikeOutlined)}
				<span className="comment-action">{like}</span>
			</span>
		</Tooltip>,
		<Tooltip key="comment-basic" title="Comment">
			<span>
				{React.createElement(MessageOutlined)}
				<span className="comment-action">{comment}</span>
			</span>
		</Tooltip>,
	];

	return (
		<Comment
			actions={actions}
			author={author.name}
			avatar={<Avatar src={author.avatar} alt="Avatar" />}
			content={
				<>
					<Typography.Paragraph>{described}</Typography.Paragraph>
					{post.image ? (
						<Image.PreviewGroup>
							<Space size={"large"} wrap>
								{post.image.map((image) => {
									return <Image key={image.id} width={300} src={image.url} />;
								})}
							</Space>
						</Image.PreviewGroup>
					) : null}
					{post.video ? (
						<div style={{ display: "flex", justifyContent: "center" }}>
							<video width="600" poster={post.video.thumb} controls>
								<source src={post.video.url} type="video/mp4" />
							</video>
						</div>
					) : null}
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
