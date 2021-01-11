package main

// import "github.com/ogunb/matchday-functions/fixture/queue"

import "github.com/ogunb/matchday-functions/fixture"

func main() {
	// queue.AddToCloudQueue()
	fixture.FetchFixtures(nil, fixture.PubSubMessage{})
}