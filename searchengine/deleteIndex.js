const client = require('./config/connection');

const deleteIndex = async (index) => {
  const { body } = await client.indices.exists({ index: index });
  if (!body) return await client.indices.delete({ index: index });
  return;
};

(async function () {
  try {
    await deleteIndex('posts');
  } catch (err) {
    console.log(err);
  }
})();
