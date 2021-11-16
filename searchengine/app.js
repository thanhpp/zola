const express = require('express');
const app = express();
let port = process.env.PORT || 3000;
//const routes

app.use(express.json());

//app.use(routes)

app.listen(port, () => {
  console.log(`The server is listening on port ${port}`);
});
