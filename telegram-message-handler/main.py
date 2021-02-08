import re

from commands.start import start_handler
from commands.add_team import add_team_handler
from commands.add_team_with_id import add_team_with_id_handler

from api.telegram_api import answer_callback

COMMAND_REGEX = re.compile("^\/\S+")

def telegram_message_handler(request):
  body = request.get_json()

  print(body)

  if not "update_id" in body:
    print("Request is not from telegram.")
    return;

  callback_query = body.get("callback_query")
  if callback_query:
    message = callback_query.get("message")
    callback_text = add_team_with_id_handler(
      chat_id=message.get("chat").get("id"),
      arguments=[callback_query.get("data")]
    )
    answer_callback(callback_query.get("id"), callback_text)
    return "OK"

  message = body.get("message")
  text = message.get("text")
  command_match = re.match(COMMAND_REGEX, text)

  if not command_match:
    print("No command was provided.")
    return

  command = command_match.group()
  arguments = re.sub(COMMAND_REGEX, "", text).strip().split()

  print(command)
  print(arguments)

  message_handlers = {
    "/start": start_handler,
    "/addteam": add_team_handler,
    "/addteamwithid": add_team_with_id_handler,
    # TODO: /myteams
    # TODO: /removeteam id
  }

  handler = message_handlers.get(command.lower())

  if not handler:
    # TODO
    print("Not a valid command.")
    return

  handler(message["chat"].get("id"), arguments)

  return "OK"
