from db.firestore import remove_follower


def remove_team_with_id_handler(chat_id, arguments):
  team_id = arguments[0] if arguments and len(arguments) > 0 else ""
  remove_follower(team_id=team_id, chat_id=chat_id)

  return "Successfully removed."