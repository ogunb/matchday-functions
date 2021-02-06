import re

from commands.start import start_handler
from commands.add_team import add_team_handler

COMMAND_REGEX = re.compile("^\/\S+")

def telegram_message_handler(request):
  body = request.get_json()

  print(body)

  if not 'update_id' in body:
    print('Request is not from telegram.')
    return;

  message = body.get('message')
  text = message.get('text')
  commandMatch = re.match(COMMAND_REGEX, text)

  if not commandMatch:
    print('No command was provided.')
    return

  command = commandMatch.group()
  arguments = re.sub(COMMAND_REGEX, "", text).strip()

  print(command)
  print(arguments)

  message_handlers = {
    "/start": start_handler,
    "/addteam": add_team_handler,
  }

  handler = message_handlers.get(command.lower())

  if not handler:
    # TODO
    print('Not a valid command.')
    return

  handler(message["chat"].get('id'), arguments)

  return 'OK'
