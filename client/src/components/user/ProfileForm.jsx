import React, { useEffect } from "react";
import "antd/dist/antd.css";
import { Card, Form, Input, Button, Space } from "antd";

//import { UploadOutlined } from "@ant-design/icons";
//import AuthContext from "../../context/authContext";

export default function ProfileForm(props) {
	//console.log(props.user);

	const {
		//id,
		username,
		phone,
		description,
		// avatar,
		// cover_image,
		link,
		address,
		city,
		country,
		name,
	} = props.user;
	// const [files, setFiles] = useState({
	// 	avatar: avatar,
	// 	cover_img: cover_img,
	// });
	const [form] = Form.useForm();

	//clean up form values and validate
	const onFinish = (values) => {
		//create formData
		const formData = new FormData();
		for (const property in values) {
			formData.append(`${property}`, values[`${property}`]);
		}
		// for (var value of formData.values()) {
		// 	console.log(value);
		// }
		props.editUserHandler(formData);
	};

	useEffect(() => {
		form.resetFields();
	}, [props.user, form]);

	//get file name?
	// const normFile = (e) => {
	// 	console.log("Upload event:", e, e.fileList);
	// 	//e.fileList[0].originFileObj
	// 	if (Array.isArray(e)) {
	// 		return e;
	// 	}

	// 	return e && e.fileList;
	// };

	return (
		<Card>
			<Form
				onFinish={onFinish}
				form={form}
				layout="vertical"
				name="profile_form"
				initialValues={{
					name: name,
					username: username,
					phone: phone,
					link: link,
					address: address,
					city: city,
					country: country,
					description: description,
				}}
			>
				<Form.Item label="Name" name="name">
					<Input />
				</Form.Item>
				<Form.Item label="Username" name="username">
					<Input />
				</Form.Item>
				<Form.Item label="Phone number" name="phone">
					<Input disabled />
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
				<Form.Item name="link" label="Website">
					<Input />
				</Form.Item>
				<Form.Item name="description" label="Description">
					<Input.TextArea maxLength={150} />
				</Form.Item>
				{/* <Form.Item
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
						beforeUpload={() => {
							return false;
						}}
					>
						<Button icon={<UploadOutlined />}>Click to upload</Button>
					</Upload>
				</Form.Item>
				<Form.Item
					name="cover_image"
					label="Cover Image"
					valuePropName="fileList"
					getValueFromEvent={normFile}
				>
					<Upload
						maxCount={1}
						name="logo"
						listType="picture"
						beforeUpload={() => {
							return false;
						}}
					>
						<Button icon={<UploadOutlined />}>Click to upload</Button>
					</Upload>
				</Form.Item> */}
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
