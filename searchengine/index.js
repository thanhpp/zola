const client = require('./config/connection');
const axiosClient = require('./config/axiosClient');

const createIndex = async (index) => {
  const { body } = await client.indices.exists({ index: index });
  if (!body) {
    await client.indices.create({ index });
  }
};

const createUserIndex = async () => {
  try {
    const { body } = await client.indices.exists({ index: 'users' });
    if (!body) {
      await client.indices.create(
        {
          index: 'users',
          // body: {
          //   mappings: {
          //     properties: {
          //       id: { type: 'keyword' },
          //       name: { type: 'text' },
          //       username: { type: 'text' },
          //       phonenumber: { type: 'text' },
          //     },
          //   },
          // },
        },
        { ignore: [400] },
      );
    }
    const response = await axiosClient.get('/admin/users');
    const docs = response.data.data.users.flatMap((user) => {
      let userObj = {
        id: user.user_id,
        name: user.name,
        username: user.username,
        phonenumber: user.phone,
      };
      return [{ index: { _index: 'users', _id: user.user_id } }, userObj];
    });
    // console.log(docs);
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
    const { body: count } = await client.count({ index: 'users' });
    console.log(count);
  } catch (err) {
    console.log(err);
  }
};

(async function () {
  try {
    //const p1 = createIndex('post');
    await createUserIndex();
    // await client.indices.delete({
    //   index: 'users',
    // });
    //const p3 = createIndex('message');
    //await Promise.all([p1, p2, p3]);
  } catch (err) {
    console.log(err);
  }
})();
