const client = require('../config/connection');
const index = 'message';

async function getMessages() {
  const query = {
    query: {
      match_all: {},
    },
    size: 10000,
  };
  return await client.search({ index, body: query });
}

async function addManyMessages(messages) {
  const docs = messages.flatMap((message) => {
    let messageObj = {
      id: message.user_id,
      name: message.name,
      username: message.username,
      phonenumber: message.phone,
    };
    return [{ index: { _index: index, _id: message.user_id } }, messageObj];
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

async function searchMessages(searchItem) {
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

async function getMessage(id) {
  return await client.get({ index, id });
}

async function deleteMessage(id) {
  return await client.delete({ index, id });
}

async function addMessage({ id, ...message }) {
  return await client.index({ index, id, body: message });
}

module.exports = {
  getMessage,
  getMessages,
  deleteMessage,
  addMessage,
  searchMessages,
  addManyMessages,
};
