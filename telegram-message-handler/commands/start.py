from api.telegramApi import send_message

def start_handler(chat_id, arguments):
  # TODO read from .md file.
  text="*Welcome to Matchday Reminder* \n Here is a list of available commands: \n `/addTeam {teamName}` Search and select a team to be notified about _\\(team name is required\\)_"
  json = send_message(chat_id=chat_id, text=text)
  print(json)
