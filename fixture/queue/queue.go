package queue

import (
	"log"

	"github.com/ogunb/matchday-functions/fixture/model"
)

func PurgeCloudQueue() {
	log.Println("TODO: PURGE CLOUD QUEUE")
}

func AddToCloudQueue(match model.Match) {
	log.Println(match)
}
