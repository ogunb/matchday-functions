package fixture

import (
	"fmt"
	"net/http"

	"github.com/ogunb/matchday-functions/fixture/apis"
)

var sportsAPI = apis.NewSportsService()

// FetchFixtures fetches next 5 events for given team
func FetchFixtures(w http.ResponseWriter, r *http.Request) {
	fixture := sportsAPI.FetchNotStartedMatches("549")

	fmt.Printf("%v", fixture)
	//fixture := api.FetchNotStartedMatches()
	//
	//if len(fixture) == 0 {
	//	log.Fatal("No event was found.")
	//}
	//
	//queue.PurgeQueue()
	//
	//for _, fixture := range fixture {
	//	if fixture.Status == "Not Started" {
	//		queue.CreateTask(fixture)
	//	}
	//}
	//
	//return nil
}
