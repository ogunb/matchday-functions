from api.telegram_api import send_message

def add_team_with_id_handler(chat_id, arguments):
  team_id = arguments[0] if arguments and len(arguments) > 0 else None

  if not team_id:
    print("No team id was provided.")
    send_message(chat_id=chat_id, text="Something went wrong: `Team id` is required.")
    return;

  print(chat_id)
