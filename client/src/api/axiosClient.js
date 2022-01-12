const axios = require("axios");

const axiosClient = axios.create({
	baseURL: process.env.REACT_APP_API_URL,
});

axiosClient.interceptors.request.use((config) => {
	config.headers["Authorization"] = `Bearer ${localStorage.getItem("token")}`;
	config.headers["content-type"] = "multipart/form-data";
	return config;
});

axiosClient.interceptors.response.use(
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

module.exports = axiosClient;
