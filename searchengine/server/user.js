const client = require('../config/connection');
const index = 'users';

async function addUser(user) {
  let userObj = {
    id: user.user_id,
    name: user.name,
    username: user.username,
    phonenumber: user.phone,
  };
  return await client.index({
    index,
    id: user.id,
    refresh: 'true',
    body: userObj,
  });
}

async function addManyUsers(users) {
  const docs = users.flatMap((user) => {
    let userObj = {
      id: user.user_id,
      name: user.name,
      username: user.username,
      phonenumber: user.phone,
    };
    return [{ index: { _index: index, _id: user.user_id } }, userObj];
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

async function editUser({ id, ...editedUser }) {
  return await client.update({
    index,
    refresh: 'true',
    id: id,
    body: {
      doc: editedUser,
    },
  });
}

async function deleteUser(id) {
  return await client.delete({ index, refresh: 'true', id });
}

async function getUser(id) {
  return await client.get({ index, id });
}

async function searchUsers(searchItem) {
  const query = {
    query: {
      query_string: {
        query: `*${searchItem.keyword}*`,
        fields: ['name', 'username', 'phonenumber'],
      },
    },
    from: searchItem.index,
    size: searchItem.count,
  };
  return await client.search({ index, body: query });
}

async function getUsers() {
  const query = {
    query: {
      match_all: {},
    },
    size: 10000,
  };
  return await client.search({ index, body: query });
}

module.exports = {
  addUser,
  editUser,
  deleteUser,
  getUser,
  getUsers,
  addManyUsers,
  searchUsers,
};
