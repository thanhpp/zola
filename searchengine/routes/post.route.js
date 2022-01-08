const express = require('express');
const router = express.Router();
const post = require('../controller/post.controller');

router.post('/post', post.add);

router.post('/posts', post.addMany);

router.get('/posts', post.getAll);

router.put('/post/:id', post.update);

router.get('/post/:id', post.getOne);

router.post('/search/post', post.search);

router.delete('/post/:id', post.delete);

router.use((err, req, res, next) => {
  console.log(err);
  const { code = 500, message = 'Oops,something went wrong' } = err;
  return res.status(code).send(message);
});

module.exports = router;
