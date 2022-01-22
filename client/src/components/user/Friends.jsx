import React from "react";
import "antd/dist/antd.css";
import InfiniteScroll from "react-infinite-scroll-component";
import { List, Avatar, Divider, Typography } from "antd";
import { Link } from "react-router-dom";
import { UserOutlined } from "@ant-design/icons";
const { Title } = Typography;

export default function Friends(props) {
	const { hasNextPage, fetchNextPage, total, page } = props;

	return (
		<div
			id="scrollableDiv"
			style={{
				height: "47vh",
				overflow: "auto",
				padding: "0 16px",
				border: "1px solid rgba(140, 140, 140, 0.35)",
			}}
		>
			<InfiniteScroll
				dataLength={10}
				next={fetchNextPage}
				hasMore={hasNextPage}
				endMessage={<Divider plain />}
				scrollableTarget="scrollableDiv"
			>
				<List
					dataSource={page}
					header={<Title level={4}>Friend List, total: {total}</Title>}
					renderItem={(items) =>
						items.data.data.friends
							? items.data.data.friends.map((item) => {
									return (
										<List.Item key={item.user_id}>
											<List.Item.Meta
												avatar={
													item.avatar ? (
														<Avatar alt="avatar" src={item.avatar} />
													) : (
														<Avatar size="large" icon={<UserOutlined />} />
													)
												}
												title={
													<Link to={`/users/${item.user_id}`}>{item.name}</Link>
												}
												description={item.status}
											/>
										</List.Item>
									);
							  })
							: null
					}
				/>
			</InfiniteScroll>
		</div>
	);
}
