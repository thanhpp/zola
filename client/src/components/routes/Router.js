import React, { useContext } from "react";
import { Route, Routes } from "react-router-dom";
import AdminLayout from "../../containers/Layout/AdminLayout";
import UserList from "../../containers/List/UserList";
import UserDetail from "../../views/UserDetail";
import PostsList from "../../containers/List/PostsList";
import PostDetail from "../../views/PostDetail";
import ConversationsList from "../../containers/List/ConversationsList";
import Chat from "../../containers/Chat/Chat";
import Search from "../../views/Search";
import AuthContext from "../../context/authContext";
import Login from "../../views/Login";

export default function Router() {
	const authCtx = useContext(AuthContext);
	let isLogin = authCtx.user.isLogin;

	return (
		<Routes>
			{isLogin ? (
				<Route path="/" element={<AdminLayout />}>
					<Route index element={<UserList />} />
					<Route path="users" element={<UserList />} />
					<Route path="users/:id" element={<UserDetail />} />
					<Route path="posts" element={<PostsList />} />
					<Route path="posts/:id" element={<PostDetail />} />
					<Route path="messages" element={<ConversationsList />} />
					<Route path="messages/:id" element={<Chat />} />
					<Route path="search" element={<Search />} />
				</Route>
			) : (
				<Route path="/" element={<Login />} />
			)}
		</Routes>
	);
}
