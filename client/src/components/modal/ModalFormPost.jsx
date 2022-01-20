import React from "react";
import "antd/dist/antd.css";
import { Modal, Form, Input, Button, Upload, message } from "antd";
import { UploadOutlined } from "@ant-design/icons";

const beforeImageUpload = (file) => {
	const isJpgOrPng = file.type === "image/jpeg" || file.type === "image/png";
	if (!isJpgOrPng) {
		message.error("You can only upload JPG/PNG file!");
	}
	const isLt2M = file.size / 1024 / 1024 < 2;
	if (!isLt2M) {
		message.error("Image must smaller than 2MB!");
	}
	return false;
};

const beforeVideoUpload = (file) => {
	const isLt10M = file.size / 1024 / 1024 < 10;
	if (!isLt10M) {
		message.error("Video must smaller than 10MB!");
	}
	return false;
};

export default function ModalFormPost(props) {
	const { visible, onCreate, setVisible } = props;

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
						//console.log(values);
						const { described, image, video } = values;
						const formData = new FormData();
						formData.append("described", described);
						if (video) {
							formData.append("video", video.fileList[0].originFileObj);
						}
						if (image) {
							for (let i = 0; i < image.fileList.length; i++) {
								formData.append("image", image.fileList[i].originFileObj);
							}
						}
						onCreate(formData);
						form.resetFields();
						setVisible(!visible);
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
				<Form.Item name="described" label="Write your thought here">
					<Input.TextArea maxLength={500} />
				</Form.Item>
				<Form.Item name="image" label="Attachment" valuePropName="file">
					<Upload
						maxCount={4}
						listType="picture"
						beforeUpload={beforeImageUpload}
						accept="image/*"
					>
						<Button icon={<UploadOutlined />}>Images</Button>
					</Upload>
				</Form.Item>
				<Form.Item name="video" label="Attachment" valuePropName="file">
					<Upload
						accept="video/*"
						maxCount={1}
						beforeUpload={beforeVideoUpload}
					>
						<Button icon={<UploadOutlined />}>Video</Button>
					</Upload>
				</Form.Item>
			</Form>
		</Modal>
	);
}
