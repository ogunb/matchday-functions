from db.firestore import get_user_teams

def my_teams_handler(chat_id):
  teams = get_user_teams(chat_id)

