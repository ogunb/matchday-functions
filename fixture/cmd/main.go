package main

import "github.com/ogunb/matchday-functions/fixture"

func main() {
	fixture.FetchFixtures(nil, fixture.PubSubMessage{})
}