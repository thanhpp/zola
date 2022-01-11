import React, { useContext } from "react";
import "antd/dist/antd.css";
import styles from "./Sidebar.module.css";
import { Layout, Menu } from "antd";
import { sider } from "./SideBar";
import { Link, useLocation } from "react-router-dom";
import AuthContext from "../../../context/authContext";

const { Sider } = Layout;

export default function Sidebar() {
	const location = useLocation();
	const authCtx = useContext(AuthContext);
	let userId = authCtx.user.userId;
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
				<Menu.Item key={`/users/${userId}`}>
					<Link to={`/users/${userId}`}>Profile</Link>
				</Menu.Item>
			</Menu>
		</Sider>
	);
}
