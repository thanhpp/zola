import React, { useState, useEffect } from "react";
import "antd/dist/antd.css";
import InfiniteScroll from "react-infinite-scroll-component";
import { List, Avatar, Skeleton, Divider, Typography } from "antd";
import { Link } from "react-router-dom";
const { Title } = Typography;

export default function Friends({ data }) {
	//const [loading, setLoading] = useState(false);
	const { total, friends } = data;
	const [userFriends, setUserFriends] = useState(friends);

	// useEffect(() => {
	// 	const friend = {
	// 		user_id: "sdfasdf",
	// 		avatar: "https://joeschmoe.io/api/v1/random",
	// 		username: "Hey Dowey",
	// 	};
	// 	setUserFriends([...userFriends, friend]);
	// }, []);

	return (
		<div
			id="scrollableDiv"
			style={{
				height: "60vh",
				overflow: "auto",
				padding: "0 16px",
				border: "1px solid rgba(140, 140, 140, 0.35)",
			}}
		>
			<InfiniteScroll
				dataLength={userFriends.length}
				next={() => console.log("next")}
				hasMore={userFriends.length < total}
				loader={<Skeleton avatar paragraph={{ rows: 1 }} active />}
				endMessage={<Divider plain />}
				scrollableTarget="scrollableDiv"
			>
				<List
					dataSource={userFriends}
					header={<Title level={4}>Friend List, total: {total}</Title>}
					renderItem={(friend) => (
						<List.Item key={friend.user_id}>
							<List.Item.Meta
								avatar={<Avatar alt="avatar" src={friend.avatar} />}
								title={
									<Link to={`/users/${friend.user_id}`}>{friend.username}</Link>
								}
							/>
						</List.Item>
					)}
				/>
			</InfiniteScroll>
		</div>
	);
}
