const axios = require("axios");

// const axiosClient = axios.create({
// 	baseURL: process.env.REACT_APP_API_URL,
// 	headers: {
// 		Authorization: `Bearer ${localStorage.getItem("token")}`,
// 	},
// });

const axiosClient = axios.create({
	baseURL: process.env.REACT_APP_API_URL,
	// headers: {
	// 	Authorization: `Bearer ${localStorage.getItem("token")}`,
	// },
});

axiosClient.interceptors.request.use((config) => {
	config.headers["Authorization"] = `Bearer ${localStorage.getItem("token")}`;
	config.headers["content-type"] = "multipart/form-data";
	return config;
});

module.exports = axiosClient;
