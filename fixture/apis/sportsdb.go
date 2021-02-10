package api

import (
	"encoding/json"
	"github.com/ogunb/matchday-functions/fixture/model"
	"io/ioutil"
	"log"
	"net/http"
	"fmt"
)

const sportsURL = "https://www.thesportsdb.com/api/v1/json/1"
const fixtureEndpoint = "eventsnext.php"
const teamID = "133794"

func generateURL() string {
	return fmt.Sprintf("%s/%s?id=%s", sportsURL, fixtureEndpoint, teamID)
}

func fetchNextFiveMatches() []model.Match {
	var sportsClient = http.Client{}

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

	fixture := model.EventsResponse{}
	jsonErr := json.Unmarshal(body, &fixture)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	return fixture.Events
}