import requests
from requests.exceptions import HTTPError

import os

TWITTER_AUTH_TOKEN = os.getenv("TWITTER_AUTH_TOKEN")
TOTALLY_LEGIT_TWITTER_USER = os.getenv("TOTALLY_LEGIT_TWITTER_USER")

BASE_URL = f"https://api.twitter.com/2"
FETCH_USER_URL = f"{BASE_URL}/users/by/username"

def fetch_totally_legit_match_url():
    try:
        params = { "user.fields": "url" }
        headers = { "Authorization": f"Bearer {TWITTER_AUTH_TOKEN}" }
        response = requests.get(f"{FETCH_USER_URL}/{TOTALLY_LEGIT_TWITTER_USER}", params=params, headers=headers)
        response.raise_for_status()

        res = response.json()

        return res.get("data").get("url")
    except HTTPError as http_err:
        print(f"HTTP error occurred: {http_err}")
        return "Failed to fetch match url."
    except Exception as err:
        print(f"Other error occurred: {err}")
        return "Failed to fetch match url."
