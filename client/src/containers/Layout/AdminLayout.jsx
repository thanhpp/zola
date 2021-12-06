import React from "react";
import "antd/dist/antd.css";
import { Layout } from "antd";
import HeaderMenu from "../../components/layout/Header/HeaderMenu";
import Sidebar from "../../components/layout/Siderbar/Sidebar.jsx";
import { Outlet } from "react-router";

const { Content } = Layout;

export default function AdminLayout() {
	return (
		<Layout style={{ minHeight: "100vh" }}>
			<HeaderMenu />
			<Layout>
				<Sidebar />
				<Layout style={{ padding: "0 24px 24px " }}>
					<Content
						className="site-layout-background"
						style={{
							padding: 24,
							margin: 0,
							minHeight: 280,
							backgroundColor: "#fff",
						}}
					>
						<Outlet />
					</Content>
				</Layout>
			</Layout>
		</Layout>
	);
}
