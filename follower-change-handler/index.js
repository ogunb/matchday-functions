const fetch = require('node-fetch');

exports.handleFollowerChange = (event, context) => {
  const resource = context.resource;
  console.log('Function triggered by change to: ' +  resource);

  const oldFollowers = event.oldValue.fields.followers.arrayValue.values
  const currentFollowers = event.value.fields.followers.arrayValue.values

  const hadZeroFollowers = !oldFollowers || oldFollowers.length === 0
  const hasMoreThanOneFollowers = !!currentFollowers && currentFollowers.length >= 1

  if (hadZeroFollowers && hasMoreThanOneFollowers) {
    const id = Number(event.value.fields.metadata.mapValue.fields.id.integerValue);
    const name = event.value.fields.metadata.mapValue.fields.name.stringValue;

    console.log({ id, name })

    fetch(`${process.env.PROJECT_URL}/team-fixture-handler`, {
      method: 'POST',
      body: JSON.stringify({
        id,
        name,
      })
    })
  }
};
