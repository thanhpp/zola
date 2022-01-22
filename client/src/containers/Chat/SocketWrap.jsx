import React from "react";
import { SocketContextProvider } from "../../context/socketContext";

export default function SocketWrapper(props) {
	return <SocketContextProvider>{props.children}</SocketContextProvider>;
}
