import React from "react";
import "antd/dist/antd.css";
import { Form, Input, Button } from "antd";

export default function LoginForm(props) {
	const { handleLoginSubmit } = props;

	const onFinishFailed = (errorInfo) => {
		console.log("Failed:", errorInfo);
	};

	return (
		<Form
			name="basic"
			labelCol={{
				span: 8,
			}}
			wrapperCol={{
				span: 16,
			}}
			initialValues={{
				phonenumber: "0965508091",
				password: "Ppt190898",
			}}
			onFinish={handleLoginSubmit}
			onFinishFailed={onFinishFailed}
			autoComplete="off"
		>
			<Form.Item
				label="Phone Number"
				name="phonenumber"
				rules={[
					{
						required: true,
						message: "Please input your phone number!",
					},
				]}
			>
				<Input />
			</Form.Item>

			<Form.Item
				label="Password"
				name="password"
				rules={[
					{
						required: true,
						message: "Please input your password!",
					},
				]}
			>
				<Input.Password />
			</Form.Item>

			<Form.Item
				wrapperCol={{
					offset: 8,
					span: 16,
				}}
			>
				<Button type="primary" htmlType="submit">
					Submit
				</Button>
			</Form.Item>
		</Form>
	);
}
