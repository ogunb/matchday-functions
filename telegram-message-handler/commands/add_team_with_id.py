from api.telegram_api import send_message
from db.firestore import add_follower, get_team_metadata

def add_team_with_id_handler(chat_id, arguments):
  team_id = arguments[0] if arguments and len(arguments) > 0 else None

  if not team_id:
    error_text = "Something went wrong: `Team id` is required."
    print(error_text)
    send_message(chat_id=chat_id, text=error_text)
    return error_text;

  try:
    add_follower(team_id=team_id, chat_id=chat_id)
  except Exception as err:
    print(err)
    send_message(chat_id=chat_id, text="Something went wrong.")
    return "Something went wrong."
  else:
    team = get_team_metadata(team_id)
    success_message = f"Successfully added you to followers of {team.get('name')}"
    print(success_message)
    send_message(chat_id=chat_id, text=success_message)
    return success_message
