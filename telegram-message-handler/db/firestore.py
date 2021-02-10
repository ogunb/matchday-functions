from google.cloud import firestore
from slugify import slugify

from models.user import generate_user

from api.sports_api import fetch_team_with_id

db = firestore.Client()


def get_team_metadata(team_id):
    print(f"Getting {team_id}...")
    team_document_ref = db.collection("teams").document(str(team_id))
    teamDTO = team_document_ref.get(["metadata"]).to_dict()
    team = teamDTO.get("metadata") if teamDTO else None

    if not team:
        team = fetch_team_with_id(team_id)
        update_team_metadata(team)

    return team


def query_teams(name):
    print(f"Query for {name}...")
    team_document_ref = db.collection("teams").limit(1).where(
        u"metadata.name_slug", u"==", slugify(name))

    teams_stream = team_document_ref.stream()

    def get_team_metadata(team):
        return team.to_dict().get("metadata")

    teams = list(map(get_team_metadata, teams_stream))

    return teams


def update_team_metadata(team):
    team_document_ref = db.collection("teams").document(str(team.get("id")))
    team_document_ref.set(merge=True,
                          document_data={
                              u"metadata": {
                                  **team, "name_slug": slugify(team["name"])
                              }
                          })


def add_follower(team_id, chat_id):
    user = generate_user(chat_id)

    print(f"Adding follower {user} to {team_id}.")

    team_document_ref = db.collection("teams").document(str(team_id))
    team_document_ref.set(
        merge=True, document_data={u"followers": firestore.ArrayUnion([user])})


def remove_follower(team_id, chat_id):
    print(f"Removing follower with {chat_id} from {team_id}.")
    team_document_ref = db.collection("teams").document(str(team_id))
    team_document_ref.set(merge=True,
                          document_data={
                              u"followers":
                              firestore.ArrayRemove([{
                                  "chat_id": chat_id
                              }])
                          })


def get_user_teams(chat_id):
    team_document_ref = db.collection("teams").where("followers",
                                                     "array_contains",
                                                     {"chat_id": chat_id})
    teams_stream = team_document_ref.stream()

    def get_team_metadata(team):
        return team.to_dict().get("metadata")

    teams = list(map(get_team_metadata, teams_stream))

    return teams
