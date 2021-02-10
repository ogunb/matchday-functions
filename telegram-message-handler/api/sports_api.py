import requests
import os
from requests.exceptions import HTTPError

APISPORTS_KEY = os.getenv("APISPORTS_KEY")

BASE_URL = "https://v3.football.api-sports.io/"
FETCH_TEAM_WIH_ID_URL = f"{BASE_URL}/lookupteam.php?id="


def fetch_teams(team_name=None, id=None):
    try:
        response = requests.get(f"{BASE_URL}/teams",
                                params={
                                    "id": id,
                                    "name": team_name,
                                },
                                headers={"x-apisports-key": APISPORTS_KEY})
        response.raise_for_status()

        content = response.json().get("response")

        return list(map(lambda t: t["team"], content))
    except HTTPError as http_err:
        print(f"HTTP error occurred: {http_err}")
    except Exception as err:
        print(f"Other error occurred: {err}")


def fetch_team_with_id(id):
    try:
        response = fetch_teams(id=id)
        return response[0]
    except Exception as err:
        print(f"Other error occurred: {err}")
