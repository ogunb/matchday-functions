def createTeam(team):
  return {
    "id": team["idTeam"],
    "name": team["strTeam"],
    "league": team["strLeague"],
    "formedOn": team["intFormedYear"],
  }
