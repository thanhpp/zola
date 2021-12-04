import React from "react";
import { Route, Routes } from "react-router-dom";
// import PrivateRoute from "./PrivateRoute.js";
// import AdminRoute from "./AdminRoute.js";
// import UserRoute from "./UserRoute.js";

import UserList from "../../containers/UserList/UserList";

export default function Router() {
	return (
		<Routes>
			<Route path="/users" element={<UserList />} />
		</Routes>
	);
}
