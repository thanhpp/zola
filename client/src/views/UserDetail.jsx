import React from "react";
import "antd/dist/antd.css";
import { Row, Col, Tabs, message } from "antd";
import ProfileCard from "../components/user/ProfileCard";
import ProfileForm from "../components/user/ProfileForm";
import Friends from "../components/user/Friends";
import PostsList from "../containers/List/PostsList";

import {
	useMutation,
	useQuery,
	useQueryClient,
	useInfiniteQuery,
} from "react-query";
import { useParams } from "react-router-dom";
import { getUserInfo, editUserInfo, getUserFriend } from "../api/userApi";
import Spinner from "../components/spinner/Spinner";

const { TabPane } = Tabs;

export default function UserDetail() {
	const queryClient = useQueryClient();
	const { id } = useParams();
	const { data: userInfos, isLoading } = useQuery(
		["users", id],
		() => getUserInfo(id),
		{
			onError: (error) => {
				message.error({
					content: `Code: ${error.response.data.code};
				Message: ${error.response.data.message}`,
				});
			},
		}
	);

	const { isLoading: isEditLoading, mutate: editUser } = useMutation(
		editUserInfo,
		{
			onSuccess: () => {
				//queryClient.invalidateQueries("users", `${id}`);
				//queryClient.refetchQueries(["user", id]);
				queryClient.resetQueries();
			},
			onError: (error) => {
				message.error({
					content: `Code: ${error.response.data.code};
				Message: ${error.response.data.message}`,
				});
			},
		}
	);

	const {
		data: usersFriends,
		fetchNextPage,
		hasNextPage,
	} = useInfiniteQuery("friends", getUserFriend, {
		getNextPageParam: (lastPage) => {
			//console.log(lastPage.pageParam);
			if (lastPage.data.data.friends !== null) return lastPage.nextPage;
			return undefined;
		},
		onError: (error) => {
			message.error({
				content: `Code: ${error.response.data.code};
				Message: ${error.response.data.message}`,
			});
		},
	});

	if (isLoading) return <Spinner />;

	return (
		<>
			<Row gutter={[16, 16]}>
				<Col span={18} push={6}>
					<Tabs defaultActiveKey="1" type="card" size={"middle"}>
						<TabPane tab="Presonal Info" key="1">
							<ProfileForm
								user={userInfos.data.data}
								editUserHandler={editUser}
							/>
							{isEditLoading ? <Spinner /> : null}
						</TabPane>
						<TabPane tab="Posts" key="2">
							<PostsList id={userInfos.data.data.id} />
						</TabPane>
					</Tabs>
				</Col>

				{/* display name,avatar,friend, online */}
				<Col span={6} pull={18}>
					<ProfileCard user={userInfos.data.data} />
					<div style={{ height: "1rem" }} />
					{usersFriends && (
						<Friends
							hasNextPage={hasNextPage}
							fetchNextPage={fetchNextPage}
							total={userInfos.data.data.listing}
							page={usersFriends.pages}
						/>
					)}
				</Col>
			</Row>
		</>
	);
}
