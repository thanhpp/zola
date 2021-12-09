import React from "react";
import "antd/dist/antd.css";
import { List, Avatar, Space, Typography, Popconfirm } from "antd";
import {
	MessageOutlined,
	LikeOutlined,
	DeleteOutlined,
} from "@ant-design/icons";
import { Link } from "react-router-dom";

const IconText = ({ icon, text }) => (
	<Space>
		{React.createElement(icon)}
		{text}
	</Space>
);

const { Paragraph } = Typography;

export default function Posts({ posts }) {
	return (
		<List
			itemLayout="vertical"
			size="large"
			dataSource={posts}
			renderItem={(item) => (
				<List.Item
					key={item.title}
					actions={[
						<IconText
							icon={LikeOutlined}
							text={item.like}
							key="list-vertical-like-o"
						/>,
						<IconText
							icon={MessageOutlined}
							text={item.comment}
							key="list-vertical-message"
						/>,
						<Popconfirm
							title="Sure to delete?"
							onConfirm={() => console.log(item.id)}
						>
							<DeleteOutlined />
							<span className="comment-action-delete"> Delete</span>
						</Popconfirm>,
					]}
					extra={
						item.media[0]?.includes(".png") ||
						item.media[0]?.includes(".jpg") ? (
							<img width={272} alt="logo" src={item.media[0]} />
						) : null
					}
				>
					<List.Item.Meta
						avatar={<Avatar src={item.avatar} />}
						title={<Link to={`${item.id}`}>{item.author}</Link>}
						description={
							item.media.length !== 0
								? `With ${item.media.length} media content(s) attached`
								: ""
						}
					/>
					<Paragraph
						ellipsis={
							true
								? {
										rows: 2,
										expandable: true,
										symbol: "more",
								  }
								: false
						}
					>
						{item.content}
					</Paragraph>
				</List.Item>
			)}
		/>
	);
}
