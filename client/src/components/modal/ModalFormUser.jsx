import React from "react";
import "antd/dist/antd.css";
import { Modal, Form, Input, Select } from "antd";
const { Option } = Select;

export default function ModalFormUser(props) {
	const { visible, onCreate, onCancel } = props;
	const [form] = Form.useForm();
	const { title } = props;
	return (
		<Modal
			visible={visible}
			title={title}
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
			></Form>
		</Modal>
	);
}
