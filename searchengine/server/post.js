const client = require('../config/connection');
const index = 'posts';

async function addPost({ id, ...post }) {
  return await client.index({ index, id: id, body: post });
}

async function addManyPosts(posts) {
  const docs = posts.flatMap((post) => {
    let postObj = {
      id: post.id,
      described: post.described,
      author: post.author.name,
    };
    return [{ index: { _index: index, _id: post.user_id } }, postObj];
  });
  const { body: bulkResponse } = await client.bulk({
    refresh: true,
    body: docs,
  });
  if (bulkResponse.errors) {
    const erroredDocuments = [];
    bulkResponse.items.forEach((action, i) => {
      const operation = Object.keys(action)[0];
      if (action[operation].error) {
        erroredDocuments.push({
          status: action[operation].status,
          error: action[operation].error,
          operation: body[i * 2],
          document: body[i * 2 + 1],
        });
      }
    });
    console.log(erroredDocuments);
  }
}

async function editPost(editedPost) {
  let postObj = {
    id: editedPost.id,
    described: editedPost.described,
    author: editedPost.author.name,
  };
  return await client.update({
    index,
    id: editedPost.id,
    body: {
      doc: postObj,
    },
  });
}

async function deletePost(id) {
  return await client.delete({ index, id });
}

async function getPost(id) {
  return await client.get({ index, id });
}

async function searchPosts(searchItem) {
  const query = {
    query: {
      query_string: {
        query: `*${searchItem.keyword}*`,
        fields: ['described'],
      },
    },
    from: searchItem.index,
    size: searchItem.count,
  };
  return await client.search({ index, body: query });
}

async function getPosts() {
  const query = {
    query: {
      match_all: {},
    },
    size: 10000,
  };
  return await client.search({ index, body: query });
}

module.exports = {
  addPost,
  editPost,
  deletePost,
  getPost,
  getPosts,
  addManyPosts,
  searchPosts,
};
