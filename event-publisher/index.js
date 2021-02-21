const { PubSub } = require('@google-cloud/pubsub');
const pubSubClient = new PubSub();

function publishEvent(req, res) {
  const { topicName, ...rest } = req.body;
  const data = JSON.stringify(rest);
  const dataBuffer = Buffer.from(data);

  console.log(`Publishing to ${topicName} with ${data}...`);
  pubSubClient.topic(topicName).publish(dataBuffer).then(messageId => {
    console.log(`Message ${messageId} published.`);
    res.sendStatus(200)
  }).catch(error => {
    console.error(`Received error while publishing: ${error.message}`);
    process.exitCode = 1;
    res.sendStatus(500)
  })
}

exports.publishEvent = publishEvent;
