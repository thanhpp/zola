const client = require('../config/connection')
const index = 'message'

async function getMessages() {
    const query
    return await client.search({ index, body:query })
}

async function getMessage(id) {
    return await client.get({ index,id })
}

async function deleteMessage(id) {
    return await client.delete({ index,id })
}

async function addMessage({id,...message}) {
    return await client.index({ index,id,body:message })
}

module.exports = {
    getMessage,
    getMessages,
    deleteMessage,
    addMessage
}