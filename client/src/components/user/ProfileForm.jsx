import React from "react";
import "antd/dist/antd.css";
import { Card, Form, Input, Button, Upload } from "antd";
import { UploadOutlined } from "@ant-design/icons";

export default function ProfileForm() {
	const [form] = Form.useForm();

	//clean up form values and validate
	const onFinish = (values) => {
		console.log(values);
	};

	//get file name?
	const normFile = (e) => {
		console.log("Upload event:", e);

		if (Array.isArray(e)) {
			return e;
		}

		return e && e.fileList;
	};
	return (
		<Card>
			<Form
				onFinish={onFinish}
				form={form}
				layout="vertical"
				name="profile_form"
				initialValues={
					{
						//something
					}
				}
			>
				<Form.Item label="Name" name="name">
					<Input />
				</Form.Item>
				<Form.Item name="website" label="Website">
					<Input />
				</Form.Item>
				<Form.Item
					label="Phone number"
					name="phone"
					rules={[
						{
							required: true,
							message: "Please input phone number!",
						},
					]}
				>
					<Input />
				</Form.Item>
				<Form.Item label="Address" name="address">
					<Input />
				</Form.Item>
				<Form.Item label="City" name="city">
					<Input />
				</Form.Item>
				<Form.Item label="Country" name="country">
					<Input />
				</Form.Item>
				<Form.Item name="description" label="Description">
					<Input.TextArea maxLength={150} />
				</Form.Item>
				<Form.Item
					name="avatar"
					label="Avatar"
					valuePropName="fileList"
					getValueFromEvent={normFile}
				>
					<Upload
						maxCount={1}
						name="logo"
						action="/upload.do"
						listType="picture"
					>
						<Button icon={<UploadOutlined />}>Click to upload</Button>
					</Upload>
				</Form.Item>{" "}
				<Form.Item
					name="cover_img"
					label="Cover Image"
					valuePropName="fileList"
					getValueFromEvent={normFile}
				>
					<Upload
						maxCount={1}
						name="logo"
						action="/upload.do"
						listType="picture"
					>
						<Button icon={<UploadOutlined />}>Click to upload</Button>
					</Upload>
				</Form.Item>
				<Form.Item>
					<Button type="primary" htmlType="submit">
						Submit
					</Button>
				</Form.Item>
			</Form>
		</Card>
	);
}
