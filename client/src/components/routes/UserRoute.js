import React from "react";
import { Navigate } from "react-router-dom";
import UserLayout from "../../containers/Layout/UserLayout";
import ConversationsList from "../../containers/List/ConversationsList";
import WrapperChat from "../../containers/Chat/WrapperChat";

export const userRoutes = [
	{
		path: "/",
		element: <UserLayout />,
		children: [
			//{ index: true, element: <ConversationsList /> },
			{ path: "messages", element: <ConversationsList /> },
			{ path: "/", element: <Navigate to="/messages" /> },
			{ path: "messages/:id", element: <WrapperChat /> },
			{ path: "*", element: <Navigate to="." /> },
		],
	},
];
