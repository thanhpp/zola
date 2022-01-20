import React, { useContext } from "react";
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
import { DeleteOutlined, UserOutlined } from "@ant-design/icons";
import dayjs from "dayjs";
import relativeTime from "dayjs/plugin/relativeTime";
import AuthContext from "../../context/authContext";
dayjs.extend(relativeTime);

export default function Comments(props) {
	const {
		comments,
		isLoading,
		onLoadMore,
		handleDeleteComment,
		postId,
		hasMoreComment,
	} = props;

	//userID
	const { user } = useContext(AuthContext);

	//console.log(comments);

	const loadMore = hasMoreComment ? (
		<div
			style={{
				textAlign: "center",
				marginTop: 12,
				height: 32,
				lineHeight: "32px",
			}}
		>
			<Button onClick={onLoadMore} loading={isLoading}>
				Load more comments
			</Button>
		</div>
	) : null;

	const handleDelete = (id_com, com) => {
		//console.log({ postId: postId, commentId: id_com });
		handleDeleteComment({ postId: postId, commentId: id_com, comment: com });
	};
	return (
		<>
			<List
				className="comment-list"
				itemLayout="horizontal"
				dataSource={comments.pages}
				loading={isLoading}
				loadMore={loadMore}
				renderItem={(page) =>
					page.data.data.map((comment) => {
						return (
							<li>
								<Skeleton avatar title={false} loading={isLoading} active>
									<Comment
										key={comment.id}
										actions={
											user.role === "admin"
												? [
														<Tooltip
															key="comment-basic-delete"
															title="Delete comment"
														>
															<span
																onClick={() =>
																	handleDelete(comment.id, comment.comment)
																}
															>
																<DeleteOutlined />
																Delete
															</span>
														</Tooltip>,
												  ]
												: null
										}
										author={comment.poster.name}
										avatar={
											comment.poster.avatar ? (
												<Avatar src={comment.poster.avatar} alt="avatar" />
											) : (
												<Avatar size="small" icon={<UserOutlined />} />
											)
										}
										content={
											<Typography.Paragraph>
												{comment.comment}
											</Typography.Paragraph>
										}
										datetime={
											<Tooltip title={dayjs().format("DD-MM-YYYY HH:mm:ss")}>
												<span>{dayjs.unix(comment.created).fromNow()}</span>
											</Tooltip>
										}
									/>
								</Skeleton>
							</li>
						);
					})
				}
			/>
		</>
	);
}
