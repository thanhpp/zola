import React from "react";
import { Route, Routes } from "react-router-dom";
import AdminLayout from "../../containers/Layout/AdminLayout";
import UserList from "../../containers/List/UserList";
import UserDetail from "../../views/UserDetail";
import PostsList from "../../containers/List/PostsList";
import PostDetail from "../../views/PostDetail";
import ConversationsList from "../../containers/List/ConversationsList";
import Chat from "../../containers/Chat/Chat";

export default function Router() {
	return (
		<Routes>
			<Route path="/" element={<AdminLayout />}>
				<Route path="users" element={<UserList />} />
				<Route path="users/:id" element={<UserDetail />} />
				<Route path="posts" element={<PostsList />} />
				<Route path="posts/:id" element={<PostDetail />} />
				<Route path="messages" element={<ConversationsList />} />
				<Route path="messages/:id" element={<Chat />} />
			</Route>
		</Routes>
	);
}
