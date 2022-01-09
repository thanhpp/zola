const post = require('../server/post');
const wrapAsync = require('../utlis/wrapAsync');

exports.add = wrapAsync(async (req, res, next) => {
  await post.addPost(req.body);
  return res.status(200).send('post added to ES');
});

exports.addMany = wrapAsync(async (req, res, next) => {
  await post.addManyPosts(req.body);
  return res.status(200).send('all posts added to ES');
});

exports.delete = wrapAsync(async (req, res, next) => {
  await post.deletePost(req.params.id);
  return res.status(200).send('deleted post from ES');
});

exports.getOne = wrapAsync(async (req, res, next) => {
  const result = await post.getPost(req.params.id);
  return res.status(200).send(result.body);
});

exports.update = wrapAsync(async (req, res, next) => {
  await post.editPost(req.body);
  return res.status(200).send('updated post in ES');
});

exports.getAll = wrapAsync(async (req, res, next) => {
  const result = await post.getPosts();
  return res.status(200).send(result.body.hits.hits);
});

exports.search = wrapAsync(async (req, res, next) => {
  const result = await post.searchPosts(req.body);
  return res.status(200).send(result.body.hits.hits);
});
