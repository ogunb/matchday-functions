package main

// import reminder "github.com/ogunb/matchday-functions/reminder"
import fixture "github.com/ogunb/matchday-functions/fixture"

func main() {
	// reminder.SmsUser(nil, reminder.PubSubMessage{ Message: []byte("ABC") });
	fixture.FetchFixtures(nil);
}