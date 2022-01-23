import React from "react";
import { SocketContextProvider } from "../../context/socketContext";
import WrapperChat from "./WrapperChat";

export default function SocketWrap() {
	return (
		<SocketContextProvider>
			<WrapperChat />
		</SocketContextProvider>
	);
}
