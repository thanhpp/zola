const client = require('../config/connection');
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
  return res.status(200).send(result.body._source);
});

exports.update = wrapAsync(async (req, res, next) => {
  await user.editUser(req.body);
  return res.status(200).send('updated user in ES');
});

exports.getAll = wrapAsync(async (req, res, next) => {
  const response = await user.getUsers();
  const result = response.body.hits.total.value
    ? response.body.hits.hits.map((hit) => {
        return hit._source;
      })
    : 'there is no user';
  return res.status(200).send(result);
});

exports.search = wrapAsync(async (req, res, next) => {
  const response = await user.searchUsers(req.body);
  const result = response.body.hits.total.value
    ? response.body.hits.hits.map((hit) => {
        return hit._source;
      })
    : 'no result was found';
  return res.status(200).send(result);
});
