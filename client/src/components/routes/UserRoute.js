import React from "react";
import { Navigate } from "react-router-dom";
import UserLayout from "../../containers/Layout/UserLayout";
import ConversationsList from "../../containers/List/ConversationsList";
import SocketWrap from "../../containers/Chat/SocketWrap";

export const userRoutes = [
	{
		path: "/",
		element: <UserLayout />,
		children: [
			//{ index: true, element: <ConversationsList /> },
			{ path: "messages", element: <ConversationsList /> },
			{ path: "/", element: <Navigate to="/messages" /> },
			{ path: "messages/:id", element: <SocketWrap /> },
			{ path: "*", element: <Navigate to="." /> },
		],
	},
];
