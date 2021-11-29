import React from "react";
import "antd/dist/antd.css";
import { Form, Input, Button } from "antd";
import { SendOutlined } from "@ant-design/icons";

const { TextArea } = Input;

const Editor = ({ onChange, onSubmit, submitting, value }) => (
	<>
		<Form.Item>
			<TextArea rows={4} onChange={onChange} value={value} />
		</Form.Item>
		<Form.Item>
			<Button
				htmlType="submit"
				loading={submitting}
				onClick={onSubmit}
				type="primary"
				icon={<SendOutlined />}
			>
				Send
			</Button>
		</Form.Item>
	</>
);

export default Editor;
