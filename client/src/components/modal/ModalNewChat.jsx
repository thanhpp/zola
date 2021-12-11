import React from "react";
import "antd/dist/antd.css";
import { Modal, Form, Input } from "antd";

export default function ModalNewChat(props) {
	const { visible, setVisible } = props;
	const [form] = Form.useForm();

	return (
		<Modal
			visible={visible}
			title="New Conversation"
			okText="Send"
			cancelText="Cancel"
			onCancel={() => {
				form.resetFields();
				setVisible(!visible);
			}}
			onOk={() => {
				form
					.validateFields()
					.then((values) => {
						form.resetFields();
						console.log(values);
						//onCreate(values);
					})
					.catch((info) => {
						console.log("Validate Failed:", info);
					});
			}}
		>
			<Form form={form} layout="vertical" name="form_in_modal">
				<Form.Item
					label="User ID"
					name="userId"
					rules={[
						{
							required: true,
							message: "Please input userId!",
						},
					]}
				>
					<Input />
				</Form.Item>

				<Form.Item
					label="Message"
					name="message"
					rules={[
						{
							required: true,
							message: "Please input a message!",
						},
					]}
				>
					<Input.TextArea rows={4} />
				</Form.Item>
			</Form>
		</Modal>
	);
}
