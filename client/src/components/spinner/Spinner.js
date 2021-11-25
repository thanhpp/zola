import React from "react";
import "antd/dist/antd.css";
import spinerStyle from "Spinner.module.css";
import { Spin } from "antd";

export default function Spinner() {
	return (
		<div className={spinerStyle["container"]}>
			<Spin size="large" />
		</div>
	);
}
