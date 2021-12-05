import React from "react";
import { Route, Routes } from "react-router-dom";
// import PrivateRoute from "./PrivateRoute.js";
// import AdminRoute from "./AdminRoute.js";
// import UserRoute from "./UserRoute.js";

import AdminLayout from "../../containers/Layout/AdminLayout";
import UserList from "../../containers/List/UserList";
import UserDetail from "../../views/UserDetail";

export default function Router() {
	return (
		<Routes>
			<Route path="/" element={<AdminLayout />}>
				<Route path="users" element={<UserList />} />
				<Route path="users/:id" element={<UserDetail />} />
			</Route>
		</Routes>
	);
}
