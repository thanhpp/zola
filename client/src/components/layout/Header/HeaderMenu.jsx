import React from "react";
import "antd/dist/antd.css";
import styles from "./Header.module.css";
import { Menu, Layout } from "antd";
import { LogoutOutlined } from "@ant-design/icons";

const { Header } = Layout;
export default function HeaderMenu() {
	return (
		<Header className={styles.header}>
			<div className={styles.logo} />
			<Menu theme="dark" mode="horizontal">
				<Menu.Item
					key="1"
					icon={<LogoutOutlined />}
					style={{ marginLeft: "auto" }}
				>
					Logout
				</Menu.Item>
			</Menu>
		</Header>
	);
}
