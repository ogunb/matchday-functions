package main

import (
	"log"

	"github.com/joho/godotenv"
	// "github.com/ogunb/matchday-functions/fixture/queue"
)

import "github.com/ogunb/matchday-functions/fixture"

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// queue.CreateTask()
	fixture.FetchFixtures(nil, fixture.PubSubMessage{})
}
