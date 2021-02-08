from google.cloud import firestore

from models.user import generate_user

db = firestore.Client()

def get_team_metadata(team_id):
  print(f"Getting {team_id}...")
  team_document_ref = db.collection("teams").document(str(team_id))
  team = team_document_ref.get(["metadata"]).to_dict()
  return team.get("metadata")

def update_team_metadata(team):
  team_document_ref = db.collection("teams").document(str(team.get("id")))
  team_document_ref.set(merge=True, document_data={
    u"metadata": team
  })

def add_follower(team_id, chat_id):
  user = generate_user(chat_id)

  print(f"Adding follower {user} to {team_id}.")

  team_document_ref = db.collection("teams").document(str(team_id))
  team_document_ref.set(merge=True, document_data={
    u"followers": firestore.ArrayUnion([user])
  })

def remove_follower(team_id, chat_id):
  print(f"Removing follower with {chat_id} from {team_id}.")
  team_document_ref = db.collection("teams").document(str(team_id))
  team_document_ref.set(merge=True, document_data={
    u"followers": firestore.ArrayRemove([{ "chat_id": chat_id }])
  })
