from google.cloud import firestore

db = firestore.Client()

def create_team_document():
  print("create_team")

def update_team_metadata():
  print("team")

def add_follower(team_id, chat_id):
  has_error = False
  message = "Successfully added."

  team_document_ref = db.collection("teams").document(str(team_id))
  doc = team_document_ref.get(["followers"])
  if doc.exists:
    print(f"Document data: {doc.to_dict()}")
  else:
    print(u"No such document!")

  return { has_error, message }

def remove_follower():
  print("remove_follower")


print(add_follower(123, 123))
