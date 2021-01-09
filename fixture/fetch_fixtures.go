package fixture

import (
	"context"
	"log"
)

// PubSubMessage is the payload of a Pub/Sub event.
type PubSubMessage struct {}

func FetchFixtures(ctx context.Context, m PubSubMessage) error {
	log.Printf("Hello")
	return nil
}
