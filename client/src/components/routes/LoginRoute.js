import React from "react";
import Login from "../../views/Login";
import { Navigate } from "react-router-dom";
export const loginRoutes = [
	{ index: true, element: <Login /> },
	{ path: "/login/*", element: <Login /> },
	{ path: "*", element: <Navigate to="." /> },
];
