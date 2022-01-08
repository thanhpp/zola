const elasticsearch = require('@elastic/elasticsearch');
const client = new elasticsearch.Client({
  //node: 'http://localhost:9200',
  node: 'https://9al1wnxwvy:qyrdj8jdil@olive-688332222.us-east-1.bonsaisearch.net:443',
  log: 'error',
});

module.exports = client;
