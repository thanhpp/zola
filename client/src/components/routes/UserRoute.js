import React, { useContext } from "react";
import { Route, Redirect } from "react-router-dom";
import AuthContext from "../../context/authContext";

const accessRole = "user";

export default function AdminRoute({ children, ...rest }) {
	const authCtx = useContext(AuthContext);
	const { isLogin, role } = authCtx.user;
	const userHasPermission = role && accessRole;

	if (!isLogin) {
		return <Redirect to="/login" />;
	}

	return (
		<Route
			{...rest}
			render={() => {
				return userHasPermission === true ? (
					children
				) : (
					<Redirect to="/error/401" />
				);
			}}
		/>
	);
}
