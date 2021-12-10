import React from "react";
import "antd/dist/antd.css";
import { Modal, Form, Input } from "antd";

export default function ModalFormUser(props) {
	const { visible, onCreate, onCancel } = props;
	const [form] = Form.useForm();
	return (
		<Modal
			visible={visible}
			title="Add new user"
			okText="Save"
			cancelText="Cancel"
			onCancel={() => {
				form.resetFields();
				onCancel();
			}}
			onOk={() => {
				form
					.validateFields()
					.then((values) => {
						form.resetFields();
						onCreate(values);
					})
					.catch((info) => {
						console.log("Validate Failed:", info);
					});
			}}
		>
			<Form
				form={form}
				layout="vertical"
				name="form_in_modal"
				initialValues={
					{
						//something
					}
				}
			>
				<Form.Item
					label="Phone number"
					name="phoneNumber"
					rules={[
						{
							required: true,
							message: "Please input phone number!",
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
							message: "Please input password!",
						},
					]}
				>
					<Input.Password />
				</Form.Item>
			</Form>
		</Modal>
	);
}
