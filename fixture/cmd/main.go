package main

import (
	"github.com/ogunb/matchday-functions/fixture/model"
	"github.com/ogunb/matchday-functions/fixture/services"
)

func main() {
	// services.CreateTask()
	// fixture.FetchFixtures(nil, nil)
	//sportsAPI := apis.NewSportsService()
	//fixture := sportsAPI.FetchNotStartedMatches("549")
	//firestore := db.NewFirestoreService()
	//firestore.GetTeamsWithFollowers()
	team := model.Team{ID: 1233, Name: "1233"}
	queue := services.NewQueueService()
	queuePath := queue.GenerateQueuePath(team)
	queue.CreateQueue(queuePath)
	//fmt.Printf("%v", fixture)
}
