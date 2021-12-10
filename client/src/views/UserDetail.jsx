import React from "react";
import "antd/dist/antd.css";
import { Row, Col } from "antd";
import ProfileCard from "../components/user/ProfileCard";
import ProfileForm from "../components/user/ProfileForm";

const userInfo = {
	id: "1234412341",
	phoneNumber: "0981209471",
	username: "Jane Doe",
	description: "This is definately a legit profile",
	avatar: "https://joeschmoe.io/api/v1/random",
	cover_img:
		"https://gw.alipayobjects.com/zos/rmsportal/JiqGstEfoWAOHiTxclqi.png",
	link: "example.com/Jane-Doe",
	address: "Earth",
	city: "LA",
	country: "US",
	listing: "12",
	is_friend: "1", // 1 : isfriend, 0: is not
	online: "1", //1: online; 0: offline
	state: "1", // 0: inactive ; 1: active
};

export default function UserDetail() {
	return (
		<Row gutter={[16, 16]}>
			<Col span={18} push={6}>
				<ProfileForm user={userInfo} />
			</Col>

			{/* display name,avatar,friend, online */}
			<Col span={6} pull={18}>
				<ProfileCard
					userId={userInfo.id}
					online={userInfo.online}
					description={userInfo.description}
					username={userInfo.username}
					coverImg={userInfo.cover_img}
					avatar={userInfo.avatar}
					state={userInfo.state}
				/>
			</Col>
		</Row>
	);
}
