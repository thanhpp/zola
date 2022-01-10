import React from "react";
import "antd/dist/antd.css";
import { Row, Col, Space, Tabs } from "antd";
import ProfileCard from "../components/user/ProfileCard";
import ProfileForm from "../components/user/ProfileForm";
import Friends from "../components/user/Friends";
import PostsList from "../containers/List/PostsList";
import { useMutation, useQuery, useQueryClient } from "react-query";
import { useParams } from "react-router-dom";
import { getUserInfo } from "../api/userApi";
import Spinner from "../components/spinner/Spinner";

const { TabPane } = Tabs;

const userInfo = {
	id: "1234412341",
	phoneNumber: "0981209471",
	username: "Jane Doe",
	description: "This is definitely a legit profile",
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

const userFriendList = {
	friends: [
		{
			user_id: "1234",
			username: "joe doe",
			avatar: "https://joeschmoe.io/api/v1/random",
		},
		{
			user_id: "1235",
			username: "joe doe",
			avatar: "https://joeschmoe.io/api/v1/random",
		},
		{
			user_id: "1236",
			username: "joe doe",
			avatar: "https://joeschmoe.io/api/v1/random",
		},
		{
			user_id: "1237",
			username: "joe doe",
			avatar: "https://joeschmoe.io/api/v1/random",
		},
		{
			user_id: "1238",
			username: "joe doe",
			avatar: "https://joeschmoe.io/api/v1/random",
		},
		{
			user_id: "1239",
			username: "joe doe",
			avatar: "https://joeschmoe.io/api/v1/random",
		},
		{
			user_id: "1231",
			username: "joe doe",
			avatar: "https://joeschmoe.io/api/v1/random",
		},
		{
			user_id: "1232",
			username: "joe doe",
			avatar: "https://joeschmoe.io/api/v1/random",
		},
		{
			user_id: "1233",
			username: "joe doe",
			avatar: "https://joeschmoe.io/api/v1/random",
		},
		{
			user_id: "1212",
			username: "joe doe",
			avatar: "https://joeschmoe.io/api/v1/random",
		},
		{
			user_id: "1213",
			username: "joe doe",
			avatar: "https://joeschmoe.io/api/v1/random",
		},
	],
	total: "12",
};

export default function UserDetail() {
	const { id } = useParams();
	const { data: userInfos, isLoading } = useQuery(["users", id], () =>
		getUserInfo(id)
	);

	if (isLoading) return <Spinner />;
	//console.log(userInfos.data.data);
	return (
		<Row gutter={[16, 16]}>
			<Col span={18} push={6}>
				<Tabs defaultActiveKey="1" type="card" size={"middle"}>
					<TabPane tab="Presonal Info" key="1">
						<ProfileForm user={userInfos.data.data} />
					</TabPane>
					<TabPane tab="Posts" key="2">
						<PostsList />
					</TabPane>
				</Tabs>
			</Col>

			{/* display name,avatar,friend, online */}
			<Col span={6} pull={18}>
				<Space direction="vertical">
					<ProfileCard
						userId={userInfos.data.data.id}
						online={userInfos.data.data.online}
						description={userInfos.data.data.description}
						username={userInfos.data.data.username}
						coverImg={userInfos.data.data.cover_img}
						avatar={userInfos.data.data.avatar}
						state={userInfos.data.data.state}
					/>
					<Friends data={userFriendList} />
				</Space>
			</Col>
		</Row>
	);
}
