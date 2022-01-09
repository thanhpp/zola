import React from "react";
import { Navigate } from "react-router-dom";
import AdminLayout from "../../containers/Layout/AdminLayout";
import UserList from "../../containers/List/UserList";
import UserDetail from "../../views/UserDetail";
import PostsList from "../../containers/List/PostsList";
import PostDetail from "../../views/PostDetail";
import ConversationsList from "../../containers/List/ConversationsList";
import Chat from "../../containers/Chat/Chat";
import Search from "../../views/Search";

export const adminRoutes = [
	{
		path: "/",
		element: <AdminLayout />,
		children: [
			//{ index: true, element: <UserList /> },
			{ path: "/", element: <Navigate to="/users" /> },
			{ path: "users", element: <UserList /> },
			{ path: "users/:id", element: <UserDetail /> },
			{ path: "posts", element: <PostsList /> },
			{ path: "posts/:id", element: <PostDetail /> },
			{ path: "messages", element: <ConversationsList /> },
			{ path: "messages/:id", element: <Chat /> },
			{ path: "search", element: <Search /> },
			//{ path: "*", element: <Navigate to="." /> },
		],
	},
];
