package main

import "github.com/ogunb/matchday-functions/fixture"

func main() {
	// queue.CreateTask()
	fixture.FetchFixtures(nil, nil)
	//sportsAPI := apis.NewSportsService()
	//fixture := sportsAPI.FetchNotStartedMatches("549")
	//fmt.Printf("%v", fixture)
}
