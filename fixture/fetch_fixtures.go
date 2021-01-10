package fixture

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/ogunb/matchday-functions/fixture/queue"
)

// PubSubMessage is the payload of a Pub/Sub event.
type PubSubMessage struct {}

type Match struct {
	Event     string `json:"strEvent"`
	Timestamp string `json:"strTimestamp"`
}

type EventsResponse struct {
	Events []Match `json:"events"`
}

const sportsURL = "https://www.thesportsdb.com/api/v1/json/1"
const fixtureEndpoint = "eventsnext.php"
const teamID = "133794"

var sportsClient = http.Client{}

func generateURL() string {
	return fmt.Sprintf("%s/%s?id=%s", sportsURL, fixtureEndpoint, teamID)
}

func fetchNextFiveMatches() []Match {
	url := generateURL()

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	res, getErr := sportsClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	fixture := EventsResponse{}
	jsonErr := json.Unmarshal(body, &fixture)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	return fixture.Events
}

// FetchFixtures fetches next 5 events for given team
func FetchFixtures(ctx context.Context, m PubSubMessage) error {
	fixture := fetchNextFiveMatches()

	if len(fixture) == 0 {
		log.Fatal("No event was found.")
	}

	queue.PurgeCloudQueue()

	for _, fixture := range fixture {
		queue.AddToCloudQueue(fixture)
	}

	return nil
}
