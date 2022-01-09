const axios = require("axios");

const axiosClient = axios.create({
	baseURL: process.env.REACT_APP_API_URL,
	headers: {
		Authorization: `Bearer ${localStorage.getItem("token")}`,
	},
});

module.exports = axiosClient;
