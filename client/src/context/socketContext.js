import React, { useContext } from "react";
import AuthContext from "./authContext";

const SocketContext = React.createContext({
	socket: null,
});

export const SocketContextProvider = (props) => {
	const { user } = useContext(AuthContext);
	const { userId } = user;
	const socket = new WebSocket(
		`${process.env.REACT_APP_CHAT_URL}?id=${userId}`
	);
	const contextValue = {
		socket: socket,
	};
	return (
		<SocketContext.Provider value={contextValue}>
			{props.children}
		</SocketContext.Provider>
	);
};

export default SocketContext;
