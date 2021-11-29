import React from "react";
import "antd/dist/antd.css";
import { Layout } from "antd";
import HeaderMenu from "../../components/layout/Header/HeaderMenu";

const { Content } = Layout;

export default function UserLayout(props) {
	return (
		<Layout style={{ minHeight: "100vh" }}>
			<HeaderMenu />
			<Layout style={{ padding: "0 24px 24px" }}>
				<Content
					className="site-layout-background"
					style={{
						padding: 24,
						margin: 0,
						minHeight: 280,
					}}
				>
					{props.children}
				</Content>
			</Layout>
		</Layout>
	);
}
