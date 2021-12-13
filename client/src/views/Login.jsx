import React, { useContext } from "react";
import CardWrapper from "../components/layout/Card/CardWrapper";
import LoginForm from "../components/forms/LoginForm";
import { useMutation } from "react-query";
import { loginUser } from "../api/userAuthentication";
import Spinner from "../components/spinner/Spinner";
import AuthContext from "../context/authContext";

export default function Login() {
	const authCtx = useContext(AuthContext);
	const { isLoading, isError, error, mutate } = useMutation(loginUser, {
		onSuccess: (data) => {
			console.log(data);
			authCtx.login(data.token);
		},
	});
	const handleLoginSubmit = (values) => {
		const { phoneNumber, password } = values;
		//console.log(values);
		mutate(phoneNumber, password);
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
