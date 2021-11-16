const client = require('./config/connection');

const createIndex = async (index) => {
  const { body } = await client.indices.exists({ index: index });
  if (!body) {
    await client.indices.create({ index });
  }
};

(async function () {
  try {
    const p1 = createIndex('post');
    const p2 = createIndex('user');
    const p3 = createIndex('message');
    await Promise.all([p1, p2, p3]);
  } catch (err) {
    console.log(err);
  }
})();
