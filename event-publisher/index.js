function publishEvent(req, res) {
  let message = req.query.message || req.body.message || 'Hello World!';
  res.status(200).send(message);
}

exports.publishEvent = publishEvent;
