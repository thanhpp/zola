const client = require('../config/connection');
const index = 'posts';

async function addPost(post) {
  let postObj = {
    id: post.id,
    described: post.described,
    author: {
      id: post.author.id,
      name: post.author.name || post.author.username,
    },
    created: post.created,
    modified: post.modified,
  };
  return await client.index({ index, id: post.id, body: postObj });
}

async function addManyPosts(posts) {
  const docs = posts.flatMap((post) => {
    let postObj = {
      id: post.id,
      described: post.described,
      author: {
        id: post.author.id,
        name: post.author.name || post.author.username,
      },
      created: post.created,
      modified: post.modified,
    };
    return [{ index: { _index: index, _id: post.id } }, postObj];
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
    author: {
      id: editedPost.author.id,
      name: editedPost.author.name || editedPost.author.username,
    },
    created: editedPost.created,
    modified: editedPost.modified,
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
      bool: {
        must: [
          {
            match: {
              described: {
                query: `${searchItem.keyword}`,
                operator: 'AND',
              },
            },
          },
          {
            match_phrase: {
              described: {
                query: `${searchItem.keyword}`,
              },
            },
          },
        ],
        should: [
          {
            query_string: {
              query: `${searchItem.keyword}`,
              fields: ['described'],
            },
          },
        ],
        minimum_should_match: 1,
        boost: 1.0,
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
