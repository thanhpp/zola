const elasticsearch = require('@elastic/elasticsearch');
const client = new elasticsearch.Client({
    node: 'http://localhost:9200',
    log: 'error'
 });

module.exports = client