const client = require('../config/connection');
const index = 'post';

async function addPost({ id, ...user }) {
  return await client.index({ index, id: id, body: user });
}

async function editPost({ id, ...editedPost }) {
  return await client.update({
    index,
    id: id,
    body: {
      doc: editedPost,
    },
  });
}

async function deletePost(id) {
  return await client.delete({ index, id });
}

async function getPost(id) {
  return await client.get({ index, id });
}

async function getPosts() {
    const query
    return await client.search({ index, body:query })
}

module.exports = {
    addPost,
    editPost,
    deletePost,
    getPost,
    getPosts
};
