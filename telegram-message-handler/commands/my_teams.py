from api.telegram_api import send_message

from db.firestore import get_user_teams

from utils.string import generate_team_string
from utils.inline_keyboard_inputs import generate_telegram_inline_keyboard_inputs


def my_teams_handler(chat_id, arguments):
    teams = get_user_teams(chat_id)

    texts = list(map(generate_team_string, teams))
    text = '\n'.join(texts)
    send_message(chat_id=chat_id, text=text)
