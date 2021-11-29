import React from "react";
import "antd/dist/antd.css";
import { List, Comment } from "antd";

export default function Messages({ messages }) {
	return (
		<List
			dataSource={messages}
			itemLayout="horizontal"
			renderItem={(props) => <Comment {...props} />}
		/>
	);
}
