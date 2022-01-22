import React from "react";
import "antd/dist/antd.css";
import { Modal, Form, Select } from "antd";

const { Option } = Select;

export default function ModalNewChat(props) {
	const { visible, setVisible, onCreate, friends } = props;
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
						//console.log(values);
						onCreate(values);
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
					<Select>
						{friends.map((friend) => {
							return <Option key={friend.user_id}>{friend.user_name}</Option>;
						})}
					</Select>
				</Form.Item>
			</Form>
		</Modal>
	);
}
