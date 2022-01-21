import React from "react";
import "antd/dist/antd.css";
import { Comment, Tooltip, Avatar, Image, Space, Typography } from "antd";
import {
	LikeOutlined,
	MessageOutlined,
	LikeFilled,
	UserOutlined,
} from "@ant-design/icons";
import dayjs from "dayjs";
import relativeTime from "dayjs/plugin/relativeTime";
dayjs.extend(relativeTime);

export default function Post(props) {
	const { post, handleInteraction } = props;
	const { is_liked, like, comment, author, described, created, id } = post;
	const actions = [
		<Tooltip key="comment-basic-like" title="Like">
			<span onClick={() => handleInteraction(id)}>
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
			avatar={
				author.avatar ? (
					<Avatar src={author.avatar} alt="Avatar" />
				) : (
					<Avatar size="small" icon={<UserOutlined />} />
				)
			}
			content={
				<>
					<Typography.Paragraph>{described}</Typography.Paragraph>
					{post.images ? (
						<Image.PreviewGroup>
							<Space size={"large"} wrap>
								{post.images.map((image) => {
									return <Image key={image.id} width={300} src={image.url} />;
								})}
							</Space>
						</Image.PreviewGroup>
					) : null}
					{post.video.url ? (
						<div style={{ display: "flex", justifyContent: "center" }}>
							<video width="600" poster={post.video.thumb} controls>
								<source src={post.video.url} type="video/mp4" />
							</video>
						</div>
					) : null}
				</>
			}
			datetime={
				<Tooltip title={dayjs().format("DD-MM-YYYY HH:mm:ss")}>
					<span>{dayjs.unix(created).fromNow()}</span>
				</Tooltip>
			}
		>
			{props.children}
		</Comment>
	);
}
