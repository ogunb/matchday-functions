package fixture

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// PubSubMessage is the payload of a Pub/Sub event.
type PubSubMessage struct {}

type Fixture struct {
	Event string `json:"strEvent"`
	Timestamp string `json:"strTimestamp"`
}

type EventsResponse struct {
	Events []Fixture `json:"events"`
}

const sportsURL = "https://www.thesportsdb.com/api/v1/json/1"
const fixtureEndpoint = "eventsnext.php"
const teamID = "133794"

var sportsClient = http.Client{}

func generateURL() string {
	return fmt.Sprintf("%s/%s?id=%s", sportsURL, fixtureEndpoint, teamID);
}

func fetchNextFiveGames() []Fixture {
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

	fixtures := EventsResponse{}
	jsonErr := json.Unmarshal(body, &fixtures)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	return fixtures.Events
}

// FetchFixtures fetches next 5 events for given team
func FetchFixtures(ctx context.Context, m PubSubMessage) error {
	fixtures := fetchNextFiveGames()

	for _, fixture := range fixtures {
		fmt.Println(fixture);
	}

	return nil
}
