from api.telegram_api import send_message
from db.firestore import get_user_teams
from utils.inline_keyboard_inputs import generate_telegram_inline_keyboard_inputs
from utils.string import generate_team_string

def remove_team_handler(chat_id, arguments):
  teams = get_user_teams(chat_id)

  def generateKeyboardOptions(team):
    return generate_telegram_inline_keyboard_inputs(
      text=generate_team_string(team),
      type="/removeteamwithid",
      data=team.get("id"),
    );

  inline_keyboard_options = list(map(generateKeyboardOptions, teams))
  send_message(chat_id=chat_id, text="Select your team:", options={
    "reply_markup": { "inline_keyboard": inline_keyboard_options }
  })