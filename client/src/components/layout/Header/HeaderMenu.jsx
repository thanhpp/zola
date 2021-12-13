import React, { useContext } from "react";
import "antd/dist/antd.css";
import styles from "./Header.module.css";
import { Menu, Layout } from "antd";
import { LogoutOutlined } from "@ant-design/icons";
import { logoutUser } from "../../../api/userAuthentication";
import AuthContext from "../../../context/authContext";

const { Header } = Layout;
export default function HeaderMenu() {
	const authCtx = useContext(AuthContext);
	const handleLogoutClick = () => {
		try {
			logoutUser();
			console.log("clicked");
			authCtx.logout();
		} catch (err) {
			console.log(err);
		}
	};
	return (
		<Header className={styles.header}>
			<div className={styles.logo} />
			<Menu theme="dark" mode="horizontal">
				<Menu.Item
					key="1"
					icon={<LogoutOutlined />}
					style={{ marginLeft: "auto" }}
					onClick={handleLogoutClick}
				>
					Logout
				</Menu.Item>
			</Menu>
		</Header>
	);
}
