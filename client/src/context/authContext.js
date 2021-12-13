import React, { useState } from "react";
//import { logoutUser } from "../api/userAuthentication";

const AuthContext = React.createContext({
	user: {
		isLogin: false,
		//role: "user"
	},
	login: (token) => {},
	logout: () => {},
});

export const AuthContextProvider = (props) => {
	const token = localStorage.getItem("token");
	const [isLogin, setIsLogin] = useState(!!token);
	//const [role, setRole] = useState("");

	const loginHandler = (token) => {
		localStorage.setItem("token", token);
		setIsLogin(true);
		//setRole(role);
	};

	const logoutHandler = () => {
		localStorage.removeItem("token");
		setIsLogin(false);
		//setRole("");
	};

	const contextValue = {
		user: { isLogin },
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
