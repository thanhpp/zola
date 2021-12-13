import React, { useContext } from "react";
import { Route, Navigate } from "react-router-dom";
import AuthContext from "../../context/authContext";

export default function PrivateRoute({ children, ...rest }) {
	const authCtx = useContext(AuthContext);
	const { isLogin } = authCtx.user.isLogin;
	if (!isLogin) return <Navigate to="/" />;
	return children;
}
