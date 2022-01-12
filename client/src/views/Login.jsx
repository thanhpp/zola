import React, { useContext } from "react";
import CardWrapper from "../components/layout/Card/CardWrapper";
import LoginForm from "../components/forms/LoginForm";
import { useMutation } from "react-query";
import { loginUser } from "../api/userAuthentication";
import Spinner from "../components/spinner/Spinner";
import AuthContext from "../context/authContext";
import { useNavigate } from "react-router-dom";
import { message } from "antd";

export default function Login() {
	let navigate = useNavigate();
	const authCtx = useContext(AuthContext);
	const { isLoading, mutate } = useMutation(loginUser, {
		onSuccess: (data) => {
			//console.log(data);
			authCtx.login(data.data);
			navigate("/", { replace: true });
		},
		onError: (error) => {
			console.log(error);
			message.error({
				content: `Code: ${error.response.data.code};
				Message: ${error.response.data.message}`,
			});
		},
	});
	if (isLoading) return <Spinner />;
	const handleLoginSubmit = (values) => {
		mutate(values);
	};

	return (
		<>
			<CardWrapper>
				<LoginForm handleLoginSubmit={handleLoginSubmit} />
			</CardWrapper>
		</>
	);
}
