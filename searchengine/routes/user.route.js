const express = require('express');
const router = express.Router();
const user = require('../controller/user.controller');

router.post('/user', user.add);

router.put('/user/:id', user.update);

router.get('/user/:id', user.getOne);

router.get('/user', user.getAll);

router.delete('/user/:id', user.delete);

router.use((err, req, res, next) => {
  console.log(err);
  const { statusCode = 500, message = 'Oops,something went wrong' } = err;
  return res.status(statusCode).send(message);
  //return res.status(err.statusCode||500).send({msg:'Error',error:err||err.stack})
});

module.exports = router;
