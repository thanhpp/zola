import React, { useState, useEffect } from "react";
import jwt from "jsonwebtoken";

const AuthContext = React.createContext({
	user: null,
	login: (token) => {},
	logout: () => {},
});

const getUserInfo = (token) => {
	const { user } = jwt.decode(token);
	const { id: userId, role } = user;
	return { userId, role };
};

export const AuthContextProvider = (props) => {
	const [user, setUser] = useState(null);
	const [token, setToken] = useState(localStorage.getItem("token"));

	const loginHandler = (token) => {
		localStorage.setItem("token", token);
		setToken(token);
	};

	useEffect(() => {
		if (token) {
			const { userId, role } = getUserInfo(token);
			setUser({ userId, role });
		}
	}, [token]);

	const logoutHandler = () => {
		setUser(null);
		localStorage.removeItem("token");
	};

	const contextValue = {
		user,
		login: loginHandler,
		logout: logoutHandler,
	};

	return (
		<AuthContext.Provider value={contextValue}>
			{props.children}
		</AuthContext.Provider>
	);
};

export default AuthContext;
