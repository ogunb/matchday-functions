package main

import "github.com/ogunb/matchday-functions/fixture"

func main() {
	// services.CreateTask()
	fixture.FetchFixtures(nil, nil)
	//sportsAPI := apis.NewSportsService()
	//fixture := sportsAPI.FetchNotStartedMatches("549")
	//firestore := db.NewFirestoreService()
	//firestore.GetTeamsWithFollowers()
	// team := model.Team{ID: 1233, Name: "1233"}
	// teams := model.Teams{Home: team, Away: team}
	// fixture := model.Fixture{Timestamp: 1613679870, Status: model.Status{Long: "Not Started"}}
	// team := model.Team{ID: 333, Name: "3333"}
	// queue := services.NewQueueService()
	// queuePath := queue.GenerateQueuePath(team)
	// queue.CreateTask(queuePath, team.ID, fixture, teams)
	// queue.CreateQueue(queuePath)
	// queue.DoesQueueExist(queuePath)
	//fmt.Printf("%v", fixture)
}
