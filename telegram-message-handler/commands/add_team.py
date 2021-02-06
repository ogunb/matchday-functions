from api.telegram_api import send_message

def add_team_handler(chat_id, arguments):
  text=f"Add team message {arguments}"
  json = send_message(chat_id=chat_id, text=text)
  print(json)
