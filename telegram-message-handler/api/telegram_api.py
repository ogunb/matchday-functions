import requests
from requests.exceptions import HTTPError

import os

BOT_BASIC_AUTH = os.getenv("BOT_BASIC_AUTH")

BASE_URL = f"https://api.telegram.org/bot{BOT_BASIC_AUTH}";
SEND_MESSAGE_URL = f"{BASE_URL}/sendMessage"

def send_message(chat_id, text, options = {}):
  try:
    request = {
      "chat_id": chat_id,
      "text": text,
      "parse_mode": "Markdown",
    }

    if options.get("parse_mode"):
      request["parse_mode"] = options.get("parse_mode")
    if options.get("reply_markup"):
      request["reply_markup"] = options.get("reply_markup")

    response = requests.post(SEND_MESSAGE_URL, json=request)
    response.raise_for_status()

    res = response.json()
    print(res)
    return res
  except HTTPError as http_err:
      print(f"HTTP error occurred: {http_err}")
  except Exception as err:
      print(f"Other error occurred: {err}")
