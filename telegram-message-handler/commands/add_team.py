from db.firestore import update_team_metadata

from api.telegram_api import send_message
from api.sportsdb_api import fetch_teams

from models.team import generate_team

from utils.inline_keyboard_inputs import generate_telegram_inline_keyboard_inputs
from utils.string import generate_team_string

def add_team_handler(chat_id, arguments):
  team_name = arguments[0] if arguments and len(arguments) > 0 else ""

  if not len(team_name) > 3:
    print("Less then four letters on team name.")
    send_message(chat_id=chat_id, text="Team name cannot be less then three letters.")
    return

  print(f"Query for: {team_name}")

  teams = fetch_teams(team_name)

  if not teams or not len(teams) > 0:
    print(f"No team was found for search: {team_name}.")
    send_message(chat_id=chat_id, text=f"No team was found with name _{team_name}_.")
    return

  def get_inline_keyboard_options(t):
    team = generate_team(t)

    update_team_metadata(team);
    return generate_telegram_inline_keyboard_inputs(
      text=generate_team_string(team),
      type="/addteamwithid",
      data=team.get("id"),
    )

  inline_keyboard_options = list(map(get_inline_keyboard_options, teams))

  send_message(chat_id=chat_id, text="Select your team:", options={
    "reply_markup": { "inline_keyboard": inline_keyboard_options }
  })
