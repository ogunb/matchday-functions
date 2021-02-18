package services

import (
	"context"
	"log"
	"os"

	"cloud.google.com/go/firestore"
	"github.com/ogunb/matchday-functions/fixture/model"
	"google.golang.org/api/iterator"
)

func createClient(ctx context.Context) *firestore.Client {
	// Sets your Google Cloud Platform project ID.
	projectID := os.Getenv("PROJECT_ID")

	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	// Close client when done with
	// defer client.Close()
	return client
}

type FirestoreService struct {
	client *firestore.Client
}

func NewFirestoreService() *FirestoreService {
	ctx := context.Background()

	return &FirestoreService{
		client: createClient(ctx),
	}
}

func (s *FirestoreService) GetTeamsWithFollowers() (*[]model.Team, error) {
	ctx := context.Background()

	iter := s.client.Collection("teams").OrderBy("followers", firestore.Desc).Documents(ctx)

	var teams []model.Team

	for {
		doc, err := iter.Next()

		if err == iterator.Done {
			break
		}

		if err != nil {
			return nil, err
		}

		data := doc.Data()
		followers := data["followers"].([]interface{})

		if len(followers) == 0 {
			break
		}

		metadata := data["metadata"].(map[string]interface{})

		teams = append(teams, model.Team{
			ID: metadata["id"].(int64),
			Name: metadata["name"].(string),
		})
	}

	return &teams, nil
}
