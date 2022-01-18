import React from "react";
import "antd/dist/antd.css";
import { List, Avatar, Space, Typography, Popconfirm, Skeleton } from "antd";
import {
	MessageOutlined,
	LikeOutlined,
	DeleteOutlined,
	LikeFilled,
} from "@ant-design/icons";
import { Link } from "react-router-dom";
import InfiniteScroll from "react-infinite-scroll-component";
import dayjs from "dayjs";
import relativeTime from "dayjs/plugin/relativeTime";
dayjs.extend(relativeTime);

const IconText = ({ icon, text }) => (
	<Space>
		{React.createElement(icon)}
		{text}
	</Space>
);

const { Paragraph } = Typography;

const mediaPreview = (post) => {
	if (post.image) {
		return <img width={272} alt="images" src={post.image[0].url} />;
	} else if (post.video.url) {
		return (
			<video width={272} poster={post.video.thumb} controls>
				<source src={post.video.url} />
			</video>
		);
	} else return;
};
export default function Posts({ posts }) {
	return (
		<div
			id="scrollableDiv"
			style={{
				height: "90vh",
				overflow: "auto",
			}}
		>
			<InfiniteScroll
				next={() => console.log("next")}
				hasMore={posts.length < 50}
				loader={<Skeleton avatar paragraph={{ rows: 3 }} active />}
				scrollableTarget="scrollableDiv"
				dataLength={posts.length}
			>
				<List
					itemLayout="vertical"
					size="large"
					dataSource={posts}
					renderItem={(post) => (
						<List.Item
							key={post.title}
							actions={[
								<IconText
									icon={!!+post.is_liked ? LikeFilled : LikeOutlined}
									text={post.like}
									key="list-vertical-like-o"
								/>,
								<IconText
									icon={MessageOutlined}
									text={post.comment}
									key="list-vertical-message"
								/>,
								<Popconfirm
									title="Sure to delete?"
									onConfirm={() => console.log(post.id)}
								>
									<DeleteOutlined />
									<span className="comment-action-delete"> Delete</span>
								</Popconfirm>,
							]}
							extra={mediaPreview(post)}
						>
							<Link to={`${post.id}`}>
								<List.Item.Meta
									avatar={<Avatar src={post.author.avatar} />}
									title={post.author.username}
									description={dayjs.unix(post.created).fromNow()}
								/>
								<Paragraph
									ellipsis={{
										rows: 2,
										expandable: true,
										symbol: "more",
									}}
								>
									{post.described}
								</Paragraph>
							</Link>
						</List.Item>
					)}
				/>
			</InfiniteScroll>
		</div>
	);
}
