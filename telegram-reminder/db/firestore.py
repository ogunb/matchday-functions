from google.cloud import firestore
import sys

db = firestore.Client()

def get_team_followers(team_id):
    print(f"Getting {team_id}...")
    team_document_ref = db.collection("teams").document(str(team_id))
    teamDTO = team_document_ref.get(["followers"]).to_dict()
    followers = teamDTO.get("followers") if teamDTO else None

    if followers == None:
        print("No followers found")
        sys.exit();

    return followers
