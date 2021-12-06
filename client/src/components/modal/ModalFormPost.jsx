import React from "react";
import "antd/dist/antd.css";
import { Modal, Form, Input } from "antd";

export default function ModalFormPost(props) {
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
			<Form.Item name="content" label="Write your thought here">
				<Input.TextArea />
			</Form.Item>
		</Modal>
	);
}
