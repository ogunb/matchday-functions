package reminder

import (
	"context"
	"log"
)

// PubSubMessage is the payload of a Pub/Sub event.
type PubSubMessage struct {
	Message []byte `json:"message"`
}

// SmsUser consumes a Pub/Sub message.
func SmsUser(ctx context.Context, m PubSubMessage) error {
	name := string(m.Message) // Automatically decoded from base64.
	if name == "" {
		name = "World"
	}
	log.Printf("Hello, %s!", name)
	return nil
}
