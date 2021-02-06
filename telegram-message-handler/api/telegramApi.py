import requests
from requests.exceptions import HTTPError

import os

BOT_BASIC_AUTH = os.getenv("BOT_BASIC_AUTH")

BASE_URL = f"https://api.telegram.org/bot{BOT_BASIC_AUTH}";
SEND_MESSAGE_URL = f"{BASE_URL}/sendMessage"

def send_message(chat_id, text, options = {
  "parse_mode": "MarkdownV2"
}):
  try:
    response = requests.post(SEND_MESSAGE_URL, data={
      "chat_id": chat_id,
      "text": text,
      "parse_mode": options["parse_mode"]
    })
    response.raise_for_status()

    return response.json()
  except HTTPError as http_err:
      print(f'HTTP error occurred: {http_err}')
  except Exception as err:
      print(f'Other error occurred: {err}')
