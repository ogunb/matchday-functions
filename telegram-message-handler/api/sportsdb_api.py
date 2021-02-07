import requests
from requests.exceptions import HTTPError

BASE_URL = "https://www.thesportsdb.com/api/v1/json/1/";
FETCH_TEAMS_URL = f"{BASE_URL}/searchteams.php"

def fetch_teams(team_name):
  try:
    response = requests.get(f"{FETCH_TEAMS_URL}?t={team_name}")
    response.raise_for_status()

    return response.json().get("teams")
  except HTTPError as http_err:
    print(f"HTTP error occurred: {http_err}")
  except Exception as err:
    print(f"Other error occurred: {err}")
