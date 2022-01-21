import React from "react";
import "antd/dist/antd.css";
import { List, Avatar, Space, Typography, Popconfirm, Skeleton } from "antd";
import {
	MessageOutlined,
	LikeOutlined,
	DeleteOutlined,
	LikeFilled,
	UserOutlined,
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
		return <img width={272} alt="images" src={post.image[0]} />;
	} else if (post.video.url) {
		return (
			<video width={272} poster={post.video.thumb} controls>
				<source src={post.video.url} />
			</video>
		);
	} else return;
};
export default function Posts(props) {
	const { pages, hasNextPage, fetchNextPage, handleDelete } = props;

	//console.log(pages);

	return (
		<div
			id="scrollableDiv"
			style={{
				height: "90vh",
				overflow: "auto",
			}}
		>
			<InfiniteScroll
				next={fetchNextPage}
				hasMore={hasNextPage}
				loader={<Skeleton avatar paragraph={{ rows: 3 }} active />}
				scrollableTarget="scrollableDiv"
				dataLength={pages.length}
			>
				<List
					itemLayout="vertical"
					size="large"
					dataSource={pages}
					renderItem={(page) =>
						page.data.data.posts.map((post) => {
							return (
								<List.Item
									key={post.id}
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
											onConfirm={() => handleDelete(post.id)}
										>
											<DeleteOutlined />
											<span className="comment-action-delete"> Delete</span>
										</Popconfirm>,
									]}
									extra={mediaPreview(post)}
								>
									<Link to={`${post.id}`}>
										<List.Item.Meta
											avatar={
												post.author.avatar ? (
													<Avatar src={post.author.avatar} />
												) : (
													<Avatar icon={<UserOutlined />} />
												)
											}
											title={post.author.name}
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
							);
						})
					}
				/>
			</InfiniteScroll>
		</div>
	);
}
