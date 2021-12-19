import React, { useContext } from "react";
import CardWrapper from "../components/layout/Card/CardWrapper";
import LoginForm from "../components/forms/LoginForm";
import { useMutation } from "react-query";
import { loginUser } from "../api/userAuthentication";
import Spinner from "../components/spinner/Spinner";
import AuthContext from "../context/authContext";
import { useNavigate } from "react-router-dom";

export default function Login() {
	let navigate = useNavigate();
	const authCtx = useContext(AuthContext);
	const { isLoading, isError, error, mutate } = useMutation(loginUser, {
		onSuccess: (data) => {
			authCtx.login(data.data);
			navigate("/", { replace: true });
		},
	});
	const handleLoginSubmit = (values) => {
		mutate(values);
	};
	if (isLoading) return <Spinner />;
	if (isError) {
		console.log(error);
	}
	return (
		<CardWrapper>
			<LoginForm handleLoginSubmit={handleLoginSubmit} />
		</CardWrapper>
	);
}
