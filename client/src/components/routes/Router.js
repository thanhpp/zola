import React, { useContext, useMemo } from "react";
import { useRoutes } from "react-router-dom";
import AuthContext from "../../context/authContext";
//import Login from "../../views/Login";
import { adminRoutes } from "./AdminRoute";
import { userRoutes } from "./UserRoute";
import { loginRoutes } from "./LoginRoute";

const role = ["admin", "user"];

const checkRoutes = (user) => {
	console.log("checking routes");
	if (!user) {
		return loginRoutes;
	} else if (user && user.role.includes(role[0])) {
		return adminRoutes;
	} else if (user && user.role.includes(role[1])) {
		return userRoutes;
	}
};

export default function Router() {
	const authCtx = useContext(AuthContext);
	let user = authCtx.user;
	const routes = useMemo(() => checkRoutes(user), [user]);
	const element = useRoutes(routes);

	return <>{element}</>;
}
