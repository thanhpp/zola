import React from "react";
import "antd/dist/antd.css";
import { Card, Form, Input, Button, Upload, Space } from "antd";
import { UploadOutlined } from "@ant-design/icons";

export default function ProfileForm({ user }) {
	console.log(user);
	const {
		id,
		username,
		phoneNumber,
		description,
		avatar,
		cover_img,
		link,
		address,
		city,
		country,
	} = user;
	const [form] = Form.useForm();

	//clean up form values and validate
	const onFinish = (values) => {
		//const {avatar,cover_img} = values
		if (!values.avatar || values.avatar.length === 0) {
			values = { ...values, avatar: avatar };
		}
		if (!values.cover_img || values.cover_img.length === 0) {
			values = { ...values, cover_img: cover_img };
		}
		console.log(values, id);
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
				initialValues={{
					name: username,
					phone: phoneNumber,
					website: link,
					address: address,
					city: city,
					country: country,
					description: description,
				}}
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
						// action="/upload.do"
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
						// action="/upload.do"
						listType="picture"
					>
						<Button icon={<UploadOutlined />}>Click to upload</Button>
					</Upload>
				</Form.Item>
				<Form.Item>
					<Space>
						<Button type="primary" htmlType="submit">
							Submit
						</Button>
						<Button
							type="primary"
							htmlType="button"
							onClick={() => form.resetFields()}
						>
							Cancel
						</Button>
					</Space>
				</Form.Item>
			</Form>
		</Card>
	);
}
