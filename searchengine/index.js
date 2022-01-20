const client = require('./config/connection');

const createIndex = async (index) => {
  const { body } = await client.indices.exists({ index: index });
  if (!body) return await client.indices.create({ index: index });
  return;
};

(async function () {
  try {
    await Promise.all([createIndex('posts'), createIndex('message')]);
  } catch (err) {
    console.log(err);
  }
})();
