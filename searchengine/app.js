const express = require('express');
const app = express();
let port = process.env.PORT || 8080;
const userRouter = require('./routes/user.route');

app.use(express.json());

app.use(userRouter);

app.listen(port, () => {
  console.log(`The server is listening on port ${port}`);
});
