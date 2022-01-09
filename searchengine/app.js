const express = require('express');
const app = express();
let port = process.env.PORT || 8080;
const userRouter = require('./routes/user.route');
const postRouter = require('./routes/post.route');

app.use(express.json());

app.use(userRouter);
app.use(postRouter);

app.listen(port, () => {
  console.log(`The server is listening on port ${port}`);
});
