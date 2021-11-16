const client = require('../config/connection');
const index = 'user';

async function addUser({ id, ...user }) {
  return await client.index({ index, id: id, body: user });
}

async function editUser({ id, ...editedUser }) {
  return await client.update({
    index,
    id: id,
    body: {
      doc: editedUser,
    },
  });
}

async function deleteUser(id) {
  return await client.delete({ index, id });
}

async function getUser(id) {
  return await client.get({ index, id });
}

async function getUsers() {
    const query
    return await client.search({ index, body:query })
}

module.exports = {
  addUser,
  editUser,
  deleteUser,
  getUser,
  getUsers
};
