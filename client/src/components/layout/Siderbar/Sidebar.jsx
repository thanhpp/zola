import React from "react";
import "antd/dist/antd.css";
import styles from "./Sidebar.module.css";
import { Layout, Menu } from "antd";
import { sider } from "./SideBar";
import { Link, useLocation } from "react-router-dom";

const { Sider } = Layout;

export default function Sidebar() {
	const location = useLocation();
	return (
		<Sider width={200} className={styles["site-layout-background"]}>
			<Menu
				mode="inline"
				selectedKeys={[location.pathname]}
				style={{ height: "100%", borderRight: 0 }}
			>
				{sider.map((side) => {
					return (
						<Menu.Item key={`/${side.link}`}>
							<Link to={side.link}>{side.navName}</Link>
						</Menu.Item>
					);
				})}
			</Menu>
		</Sider>
	);
}
