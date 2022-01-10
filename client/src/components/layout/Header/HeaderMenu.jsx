import React, { useContext } from "react";
import "antd/dist/antd.css";
import styles from "./Header.module.css";
import Spinner from "../../spinner/Spinner";
import { Menu, Layout } from "antd";
import { LogoutOutlined } from "@ant-design/icons";
import { logoutUser } from "../../../api/userAuthentication";
import AuthContext from "../../../context/authContext";
import { useQuery } from "react-query";

const { Header } = Layout;
export default function HeaderMenu() {
	const { logout } = useContext(AuthContext);
	const { status, error, refetch, isLoading } = useQuery(
		"currentUser",
		logoutUser,
		{
			enabled: false,
			retry: false,
			skip: true,
			onSuccess: () => {
				logout();
			},
		}
	);

	//if (status === "loading") return <Spinner />;
	if (status === "error") {
		console.log(error);
	}
	const handleLogoutClick = () => {
		refetch();
	};

	return (
		<>
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
			{isLoading ? <Spinner /> : null}
		</>
	);
}
