const user = require('../server/user');
const wrapAsync = require('../utlis/wrapAsync');

exports.add = wrapAsync(async (req, res, next) => {
  const result = await user.addUser(req.body);
  return res.status(200).send('user added to ES');
});

exports.delete = wrapAsync(async (req, res, next) => {
  const result = await user.deleteUser(req.params.id);
  return res.status(200).send('deleted user from ES');
});

exports.getOne = wrapAsync(async (req, res, next) => {
  const result = await user.getUser(req.params.id);
  return res.status(200).send(result.body);
});

exports.update = wrapAsync(async (req, res, next) => {
  const result = await user.editUser(req.body);
  return res.status(200).send(result.body.result);
});

exports.getAll = wrapAsync(async (req, res, next) => {
  const result = await user.getUsers();
  return res.status(200).send('searching all user ...');
});
