import React, { useState } from "react";
import "antd/dist/antd.css";
import "./EditTableRow.css";
import {
	Table,
	Input,
	Select,
	Popconfirm,
	Form,
	Typography,
	Button,
} from "antd";
import ModalFormUser from "../modal/ModalFormUser";
const { Option } = Select;

const EditTableRow = (props) => {
	const [form] = Form.useForm();
	const { data, handleData, columnName, options, handleAdd, handleDelete } =
		props;
	const [editingKey, setEditingKey] = useState("");
	const [visible, setVisible] = useState(false);
	const isEditing = (record) => record.key === editingKey;

	const EditableCell = ({
		editing,
		dataIndex,
		title,
		inputType,
		record,
		index,
		children,
		...restProps
	}) => {
		let inputNode = null;
		if (dataIndex === "type") {
			inputNode = (
				<Select>
					{options.map((option) => {
						return <Option key={option.value}>{option.text}</Option>;
					})}
				</Select>
			);
		} else {
			inputNode = <Input />;
		}

		return (
			<td {...restProps}>
				{editing ? (
					<Form.Item
						name={dataIndex}
						style={{
							margin: 0,
						}}
						rules={
							dataIndex !== "note"
								? [
										{
											required: true,
											message: `không được bỏ trống ${title}!`,
										},
								  ]
								: null
						}
					>
						{inputNode}
					</Form.Item>
				) : (
					children
				)}
			</td>
		);
	};

	const edit = (record) => {
		form.setFieldsValue({
			...record,
		});
		setEditingKey(record.key);
	};

	const cancel = () => {
		setEditingKey("");
	};

	const save = async (key) => {
		try {
			const row = await form.validateFields();
			const newData = [...data];
			const index = newData.findIndex((item) => key === item.key);

			if (index > -1) {
				const item = newData[index];
				newData.splice(index, 1, { ...item, ...row });
				handleData(newData);
				setEditingKey("");
			} else {
				newData.push(row);
				handleData(newData);
				setEditingKey("");
			}
		} catch (errInfo) {
			console.log("Validate Failed:", errInfo);
		}
	};

	const editColumns = [
		{
			title: "Operation",
			dataIndex: "operation",
			render: (_, record) => {
				const editable = isEditing(record);
				return editable ? (
					<span>
						{/* eslint-disable-next-line jsx-a11y/anchor-is-valid */}
						<a
							/*eslint-disable-next-line no-script-url*/
							href="#/"
							onClick={() => save(record.key)}
							style={{
								marginRight: 8,
							}}
						>
							Save
						</a>

						<Popconfirm title="Sure to cancel?" onConfirm={cancel}>
							{/* eslint-disable-next-line jsx-a11y/anchor-is-valid */}
							<a>Cancel</a>
						</Popconfirm>
					</span>
				) : (
					<>
						<Typography.Link
							disabled={editingKey !== ""}
							onClick={() => edit(record)}
							style={{
								marginRight: 8,
							}}
						>
							Edit
						</Typography.Link>
						<Popconfirm
							title="Sure to delete?"
							onConfirm={() => handleDelete(record.key)}
						>
							{/* eslint-disable-next-line jsx-a11y/anchor-is-valid */}
							<a>Delete</a>
						</Popconfirm>
					</>
				);
			},
		},
	];

	const columns = [...columnName, ...editColumns];
	const mergedColumns = columns.map((col) => {
		if (!col.editable) {
			return col;
		}

		return {
			...col,
			onCell: (record) => ({
				record,
				inputType: col.dataIndex.includes("state") ? "select" : "text",
				dataIndex: col.dataIndex,
				title: col.title,
				editing: isEditing(record),
			}),
		};
	});

	const onCreate = (values) => {
		handleAdd(values);
		//console.log('Received values of form: ', values);
		setVisible(false);
	};

	return (
		<>
			<Button
				onClick={() => {
					setVisible(true);
				}}
				type="primary"
				style={{
					marginBottom: 16,
				}}
			>
				Add user
			</Button>
			<Form form={form} component={false}>
				<Table
					components={{
						body: {
							cell: EditableCell,
						},
					}}
					bordered
					dataSource={data}
					columns={mergedColumns}
					rowClassName="editable-row"
					pagination={{
						onChange: cancel,
					}}
				/>
			</Form>
			<ModalFormUser
				visible={visible}
				title="Yêu cầu chức năng"
				// inputsText={inputsText}
				// inputsSelect={inputsSelect}
				onCancel={() => {
					setVisible(false);
				}}
				onCreate={onCreate}
			/>
		</>
	);
};

export default EditTableRow;
