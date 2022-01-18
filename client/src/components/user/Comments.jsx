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
import { DeleteOutlined, EditOutlined, UserOutlined } from "@ant-design/icons";
import dayjs from "dayjs";
import relativeTime from "dayjs/plugin/relativeTime";
import AuthContext from "../../context/authContext";
dayjs.extend(relativeTime);

export default function Comments(props) {
	const { comments, isLoading, onLoadMore } = props;

	//userID
	const { user } = useContext(AuthContext);

	const loadMore =
		comments.length !== 0 ? (
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
				loading={isLoading}
				loadMore={loadMore}
				renderItem={(comment) => (
					<li>
						<Skeleton avatar title={false} loading={isLoading} active>
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
									// user.userId === comment.poster.id ? (
									// 	<Tooltip key="comment-basic-edit" title="Edit comment">
									// 		<EditOutlined />
									// 		Edit
									// 	</Tooltip>
									// ) : null,
								]}
								author={comment.poster.name}
								avatar={
									comment.poster.avatar ? (
										<Avatar src={comment.poster.avatar} alt="avatar" />
									) : (
										<Avatar size="small" icon={<UserOutlined />} />
									)
								}
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
