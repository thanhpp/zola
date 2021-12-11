import React from "react";
import "antd/dist/antd.css";
import { Form, Row, Col, Input, Button, Select } from "antd";

const { Option } = Select;

export default function SearchForm() {
	const [form] = Form.useForm();
	const onFinish = (values) => {
		console.log(values);
	};
	return (
		<Form
			form={form}
			name="advanced_search"
			initialValues={{
				filter: "/",
			}}
			onFinish={onFinish}
		>
			<Row gutter={24}>
				<Col span={8}>
					<Form.Item
						name="search"
						label="Search"
						rules={[
							{
								required: true,
								message: "Input something!",
							},
						]}
					>
						<Input placeholder="what do you want to search?" />
					</Form.Item>
				</Col>
				<Col span={8}>
					<Form.Item name="filter" label="Filter">
						<Select>
							<Option value="/">All</Option>
							<Option value="posts">Post</Option>
							<Option value="friends">Friends</Option>
							<Option value="messages">Messages</Option>
						</Select>
					</Form.Item>
				</Col>
			</Row>
			<Row>
				<Col
					span={24}
					style={{
						textAlign: "right",
					}}
				>
					<Button type="primary" htmlType="submit">
						Search
					</Button>
					<Button
						style={{
							margin: "0 8px",
						}}
						onClick={() => {
							form.resetFields();
						}}
					>
						Clear
					</Button>
				</Col>
			</Row>
		</Form>
	);
}
