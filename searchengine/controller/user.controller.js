const user = require('../server/user');
const wrapAsync = require('../utlis/wrapAsync');

exports.add = wrapAsync(async (req, res, next) => {
  await user.addUser(req.body);
  return res.status(200).send('user added to ES');
});

exports.addMany = wrapAsync(async (req, res, next) => {
  await user.addManyUsers(req.body);
  return res.status(200).send('all users added to ES');
});

exports.delete = wrapAsync(async (req, res, next) => {
  await user.deleteUser(req.params.id);
  return res.status(200).send('deleted user from ES');
});

exports.getOne = wrapAsync(async (req, res, next) => {
  const result = await user.getUser(req.params.id);
  return res.status(200).send(result.body);
});

exports.update = wrapAsync(async (req, res, next) => {
  await user.editUser(req.body);
  return res.status(200).send('updated user in ES');
});

exports.getAll = wrapAsync(async (req, res, next) => {
  const result = await user.getUsers();
  return res.status(200).send(result.body.hits.hits);
});

exports.search = wrapAsync(async (req, res, next) => {
  const result = await user.searchUsers(req.body);
  return res.status(200).send(result.body.hits.hits);
});
