import React, { useState } from "react";
import "antd/dist/antd.css";
import { Modal, Form, Input, Button, Space, Upload } from "antd";
import { UploadOutlined } from "@ant-design/icons";

export default function ModalFormPost(props) {
	const { visible, onCreate, setVisible } = props;
	const [isImageAttched, setIsImageAttched] = useState(false);
	const [isVideoAttched, setIsVideoAttched] = useState(false);
	const [form] = Form.useForm();
	return (
		<Modal
			visible={visible}
			title="Add a new post"
			okText="Save"
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
				name="new_post"
				initialValues={
					{
						//something
					}
				}
			>
				<Form.Item name="content" label="Write your thought here">
					<Input.TextArea />
				</Form.Item>
				<Form.Item name="media" label="Attachment">
					<Space>
						<Upload disabled={isImageAttched}>
							<Button disabled={isImageAttched} icon={<UploadOutlined />}>
								Images
							</Button>
						</Upload>
						<Upload disabled={isVideoAttched}>
							<Button disabled={isVideoAttched} icon={<UploadOutlined />}>
								Video
							</Button>
						</Upload>
					</Space>
				</Form.Item>
			</Form>
		</Modal>
	);
}
