from api.telegram_api import send_message
from api.sportsdb_api import fetch_teams

import models.team

def add_team_handler(chat_id, arguments):
  team_name = arguments[0]
  print(f"Query for: {team_name}")

  if not len(team_name) > 3:
    print("Less then four letters on team name.")
    send_message(chat_id=chat_id, text="Team name cannot be less then three letters.")
    return

  teams = fetch_teams(team_name)

  if not teams or not len(teams) > 0:
    print(f"No team was found for search: {team_name}.")
    send_message(chat_id=chat_id, text=f"No team was found with name _{team_name}_.")
    return

  def generateTelegramInlineKeyboardInputs(t):
    team = models.team.createTeam(t)

    return [{
      "text": f"{team.get('name')} ({team.get('formedOn')}) - {team.get('league')}",
      "callback_data": team.get("id")
    }]

  inlineKeyboardOptions = list(map(generateTelegramInlineKeyboardInputs, teams))

  send_message(chat_id=chat_id, text="Select your team:", options={
    "reply_markup": { "inline_keyboard": inlineKeyboardOptions }
  })
