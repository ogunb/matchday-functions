import re

from commands.start import start_handler

COMMAND_REGEX = re.compile("^\/\S+")

def telegram_message_handler(request):
  # TODO get_json
  body = {
    "update_id": "asdfiuha",
    "message": {
      "text": '/start asdfasdf'
    }
  }

  if not 'update_id' in body:
    print('Request is not from telegram.')
    return;

  text = body.get('message').get('text')
  commandMatch = re.match(COMMAND_REGEX, text)

  if not commandMatch:
    print('No command was provided.')
    return

  command = commandMatch.group()
  arguments = re.sub(COMMAND_REGEX, "", text).strip()

  print(command)
  print(arguments)

  message_handlers = {
    "/start": start_handler
  }

  handler = message_handlers.get(command)

  if not handler:
    # TODO
    print('Not a valid command.')
    return

  handler(1688953541, arguments)


telegram_message_handler({})

