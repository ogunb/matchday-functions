package main

import "github.com/ogunb/matchday-functions/fixture"

func main() {
	// queue.CreateTask()
	fixture.FetchFixtures(nil, fixture.PubSubMessage{})
}
