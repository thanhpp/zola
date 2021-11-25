import React, { useContext } from "react";
import { Route, Redirect } from "react-router-dom";
import AuthContext from "../../context/authContext";

export default function PrivateRoute({ children, ...rest }) {
	const authCtx = useContext(AuthContext);
	const { isLoggedIn } = authCtx.user;
	return (
		<Route
			{...rest}
			render={() => {
				return isLoggedIn === true ? children : <Redirect to="/login" />;
			}}
		/>
	);
}
