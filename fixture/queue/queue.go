package queue

import (
	"log"

	fixture "github.com/ogunb/matchday-functions/fixture"
)

func PurgeCloudQueue() {
	log.Println("TODO: PURGE CLOUD QUEUE")
}

func AddToCloudQueue(match fixture.Match) {
	log.Println(match)
}
