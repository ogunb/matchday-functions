def generate_team(team):
  return {
    "id": team["idTeam"],
    "name": team["strTeam"],
    "league": team["strLeague"],
    "formedAt": team["intFormedYear"],
  }
