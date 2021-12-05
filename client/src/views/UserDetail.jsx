import React from "react";
import "antd/dist/antd.css";
import { Row, Col } from "antd";
import ProfileCard from "../components/user/ProfileCard";
import ProfileForm from "../components/user/ProfileForm";

export default function UserDetail() {
	return (
		<Row gutter={[16, 16]}>
			<Col span={18} push={6}>
				<ProfileForm />
			</Col>
			<Col span={6} pull={18}>
				<ProfileCard />
			</Col>
		</Row>
	);
}
