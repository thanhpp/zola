import React from "react";
import { Route, Switch, Redirect } from "react-router-dom";
// import PrivateRoute from "./PrivateRoute.js";
// import AdminRoute from "./AdminRoute.js";
// import UserRoute from "./UserRoute.js";

export default function Router() {
	return (
		<Switch>
			{/* <Route /> */}
			{/* <PrivateRoute />
			<AdminRoute></AdminRoute>
			<UserRoute></UserRoute> */}
			<Route exact path="/">
				<Redirect to="/home" />
			</Route>
		</Switch>
	);
}
