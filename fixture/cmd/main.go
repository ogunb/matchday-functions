package main

import (
	"github.com/ogunb/matchday-functions/fixture/db"
)

func main() {
	// queue.CreateTask()
	// fixture.FetchFixtures(nil, nil)
	//sportsAPI := apis.NewSportsService()
	//fixture := sportsAPI.FetchNotStartedMatches("549")
	firestore := db.NewFirestoreService()
	firestore.GetTeamsWithFollowers()
	//fmt.Printf("%v", fixture)
}
