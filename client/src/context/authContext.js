import React, { useState } from "react";

const AuthContext = React.createContext({
	user: { isLogin: false, role: "user" },
	login: (user, role) => {},
	logout: () => {},
});

export const AuthContextProvider = (props) => {
	const [isLogin, setLogin] = useState(false);
	const [role, setRole] = useState("");

	const loginHandler = (user, role) => {
		setLogin(true);
		setRole(role);
	};

	const logoutHandler = () => {
		setLogin(false);
		setRole("");
	};

	const contextValue = {
		user: { isLogin, role },
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
