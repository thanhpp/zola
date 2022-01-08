const express = require('express');
const router = express.Router();
const user = require('../controller/user.controller');

router.post('/user', user.add);

router.post('/users', user.addMany);

router.get('/users', user.getAll);

router.put('/user/:id', user.update);

router.get('/user/:id', user.getOne);

router.post('/search/user', user.search);

router.delete('/user/:id', user.delete);

router.use((err, req, res, next) => {
  console.log(err);
  const { code = 500, message = 'Oops,something went wrong' } = err;
  return res.status(code).send(message);
});

module.exports = router;
