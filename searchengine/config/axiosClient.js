const axios = require('axios');
require('dotenv').config();

const axiosClient = axios.create({
  baseURL: process.env.END_POINT,
  headers: {
    Authorization: `Bearer ${process.env.TOKEN}`,
  },
});

module.exports = axiosClient;
