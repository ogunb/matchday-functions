function publishEvent(req, res) {
  // const {
  //   topicName,
  //   ...rest
  // } = req.body;

  console.log(req)
  // console.log(rest)

  res.status(200).send("lol");
}

exports.publishEvent = publishEvent;
