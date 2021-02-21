exports.handleFollowerChange = (event, context) => {
  const resource = context.resource;
  console.log('Function triggered by change to: ' +  resource);

  const oldFollowers = event.oldValue.fields.followers.arrayValue.values
  const currentFollowers = event.value.fields.followers.arrayValue.values

  const hadZeroFollowers = !oldFollowers || oldFollowers.length === 0
  const hasMoreThanOneFollowers = !!currentFollowers && currentFollowers.length >= 1

  if (hadZeroFollowers && hasMoreThanOneFollowers) {
    // TODO
    // Team service > create match tasks
  } else if (!hadZeroFollowers && !hasMoreThanOneFollowers) {
    // TODO
    // Team service > purge match tasks
  }
};
