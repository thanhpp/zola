import React, { useContext, useState } from "react";
import Posts from "../../components/list/Posts";
import { Button } from "antd";
import AuthContext from "../../context/authContext";
import ModalNewPost from "../../components/modal/ModalFormPost";
import { useMutation, useQuery, useQueryClient } from "react-query";
import { getPostList } from "../../api/postApi";
import Spinner from "../../components/spinner/Spinner";

const datas = {
	code: "1000",
	message: "Success",
	posts: [],
	new_items: "3",
	last_id: "23",
};
for (let i = 0; i < 23; i++) {
	datas.posts.push({
		id: i,
		image: [
			{
				id: "1",
				url: "https://gw.alipayobjects.com/zos/rmsportal/JiqGstEfoWAOHiTxclqi.png",
			},
		],
		// video: {
		// 	url: "https://archive.org/download/BigBuckBunny_124/Content/big_buck_bunny_720p_surround.mp4",
		// 	thumb:
		// 		"https://gw.alipayobjects.com/zos/rmsportal/JiqGstEfoWAOHiTxclqi.png",
		// },
		described: "post content",
		like: "124",
		comment: "124",
		is_liked: "0",
		is_blocked: "0",
		can_comment: "1",
		can_edit: "0",
		author: {
			id: "1234345",
			username: "Jane Doe",
			avatar: "https://joeschmoe.io/api/v1/random",
			online: "0",
		},
	});
}

export default function PostsList(props) {
	const queryClient = useQueryClient();
	const { user } = useContext(AuthContext);
	const [displayModal, setDisplayModal] = useState(false);
	const { data, isLoading } = useQuery("posts", getPostList);

	if (isLoading) return <Spinner />;

	return (
		<>
			<Button type="primary" block onClick={() => setDisplayModal(true)}>
				Add new post
			</Button>
			<Posts posts={data.data.posts} />
			<ModalNewPost visible={displayModal} setVisible={setDisplayModal} />
		</>
	);
}
