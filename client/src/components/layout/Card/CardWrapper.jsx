import React from "react";
import "antd/dist/antd.css";
import styles from "./CardWrapper.module.css";
import { Card } from "antd";

export default function CardWrapper(props) {
	return (
		<div className={styles["site-card-border-less-wrapper"]}>
			<div className={styles.center}>
				<Card title="Login" bordered={false} style={{ width: 500 }}>
					{props.children}
				</Card>
			</div>
		</div>
	);
}
