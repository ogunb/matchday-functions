import base64
import json
from db.firestore import get_team_followers
from api.telegram_api import send_message

def remind_match(event, context):
    pubsub_message = base64.b64decode(event['data']).decode('utf-8')
    data = json.loads(pubsub_message)

    followers = get_team_followers(data["teamId"])

    for follower in followers:
        chat_id = follower["chat_id"]
        send_message(text=data["event"], chat_id=chat_id)
