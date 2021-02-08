import re
import json

from commands.start import start_handler
from commands.add_team import add_team_handler
from commands.add_team_with_id import add_team_with_id_handler

from api.telegram_api import answer_callback

COMMAND_REGEX = re.compile("^\/\S+")

MESSAGE_HANDLERS = {
  "/start": start_handler,
  "/addteam": add_team_handler,
  "/addteamwithid": add_team_with_id_handler,
  # TODO: /myteams
  # TODO: /removeteam id
}

def telegram_message_handler(request):
  body = request.get_json()

  print(body)

  if not "update_id" in body:
    print("Request is not from telegram.")
    return;

  callback_query = body.get("callback_query")
  message = callback_query.get("message") if callback_query else body.get("message")
  text = message.get("text")
  command_match = re.match(COMMAND_REGEX, text)

  if not callback_query and not command_match:
    print("No command was provided.")
    return

  callback_data = json.loads(callback_query.get("data")) if callback_query else None

  command = callback_data.get("type") if callback_data else command_match.group()
  arguments = callback_data.get("data") if callback_data else re.sub(COMMAND_REGEX, "", text).strip().split()

  print(command)
  print(arguments)

  handler = MESSAGE_HANDLERS.get(command.lower())

  if not handler:
    # TODO
    print("Not a valid command.")
    return

  callback_text = handler(message["chat"].get("id"), arguments)
  if callback_text and callback_query:
    answer_callback(callback_query.get("id"), callback_text)

  return "OK"

add_team_handler(1688953541, ["Besiktas"])