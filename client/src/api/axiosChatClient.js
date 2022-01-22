const axios = require("axios");

const axiosChatClient = axios.create({
	baseURL: process.env.REACT_APP_CHAT_ENDPOINT,
});

axiosChatClient.interceptors.request.use((config) => {
	config.headers["Authorization"] = `Bearer ${localStorage.getItem("token")}`;
	return config;
});

axiosChatClient.interceptors.response.use(
	(res) => res,
	(error) => {
		const token = localStorage.getItem("token");
		if (error.response.status === 401 && token) {
			localStorage.removeItem("token");
			window.location.reload();
		}
		return Promise.reject(error);
	}
);

module.exports = axiosChatClient;
