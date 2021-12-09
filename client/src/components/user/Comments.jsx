import React, { useState } from "react";
import "antd/dist/antd.css";
import {
	List,
	Comment,
	Avatar,
	Tooltip,
	Typography,
	Button,
	Skeleton,
} from "antd";
import { DeleteOutlined, EditOutlined } from "@ant-design/icons";
import dayjs from "dayjs";
import relativeTime from "dayjs/plugin/relativeTime";
dayjs.extend(relativeTime);

export default function Comments({ comments }) {
	const [loading, setLoading] = useState(false);

	//userID
	const userId = "124341343";

	const onLoadMore = () => {
		setLoading(true);
		//fetching more comments
		console.log("load more comment");
		setInterval(() => {
			setLoading(false);
		}, 1500);
	};

	const loadMore = !loading ? (
		<div
			style={{
				textAlign: "center",
				marginTop: 12,
				height: 32,
				lineHeight: "32px",
			}}
		>
			<Button onClick={onLoadMore}>Load more comments</Button>
		</div>
	) : null;

	const handleDelete = (id, id_com) => {
		console.log({ id_post: id, id_com: id_com });
	};
	return (
		<>
			<List
				className="comment-list"
				itemLayout="horizontal"
				dataSource={comments}
				loading={loading}
				loadMore={loadMore}
				renderItem={(comment) => (
					<li>
						<Skeleton avatar title={false} loading={loading} active>
							<Comment
								key={comment.id}
								actions={[
									<Tooltip key="comment-basic-delete" title="Delete comment">
										<DeleteOutlined
											onClick={() =>
												handleDelete(comment.poster.id, comment.id)
											}
										/>
										Delete
									</Tooltip>,
									//edit comment
									userId === comment.poster.id ? (
										<Tooltip key="comment-basic-edit" title="Edit comment">
											<EditOutlined />
											Edit
										</Tooltip>
									) : null,
								]}
								author={comment.poster.name}
								avatar={<Avatar src={comment.poster.avatar} alt="avatar" />}
								content={
									<Typography.Paragraph>{comment.comment}</Typography.Paragraph>
								}
								datetime={
									<Tooltip title={dayjs().format("DD-MM-YYYY HH:mm:ss")}>
										<span>{dayjs.unix(comment.created).fromNow()}</span>
									</Tooltip>
								}
							/>
						</Skeleton>
					</li>
				)}
			/>
		</>
	);
}
