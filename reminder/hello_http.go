package hellohttp

import (
	"context"
	"log"
)

// PubSubMessage is the payload of a Pub/Sub event.
type PubSubMessage struct {
	Data []byte `json:"data"`
}

// HelloHTTP consumes a Pub/Sub message.
func HelloHTTP(ctx context.Context, m PubSubMessage) error {
	name := string(m.Data) // Automatically decoded from base64.
	if name == "" {
		name = "World"
	}
	log.Printf("Hello, %s!", name)
	return nil
}
