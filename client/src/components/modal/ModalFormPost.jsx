import React from "react";
import "antd/dist/antd.css";
import { Modal, Form, Input, Button, Upload } from "antd";
import { UploadOutlined } from "@ant-design/icons";

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
						console.log(values);
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
					<Input.TextArea />
				</Form.Item>
				<Form.Item name="image" label="Attachment" valuePropName="fileList">
					<Upload
						maxCount={4}
						listType="picture"
						beforeUpload={() => {
							return false;
						}}
					>
						<Button icon={<UploadOutlined />}>Images</Button>
					</Upload>
				</Form.Item>
				<Form.Item name="video" label="Attachment" valuePropName="fileList">
					<Upload
						maxCount={1}
						beforeUpload={() => {
							return false;
						}}
					>
						<Button icon={<UploadOutlined />}>Video</Button>
					</Upload>
				</Form.Item>
			</Form>
		</Modal>
	);
}
