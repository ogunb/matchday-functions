package apis

import (
	"log"
	"os"

	"github.com/dghubble/sling"
	"github.com/ogunb/matchday-functions/team-fixture-tasks/model"
)

const baseURL = "https://v3.football.api-sports.io"

type SportsService struct {
	sling *sling.Sling
}

func NewSportsService() *SportsService {
	return &SportsService{
		sling: sling.New().Base(baseURL).Set("x-apisports-key", os.Getenv("APISPORTS_KEY")),
	}
}

func (s *SportsService) FetchNotStartedMatches(teamID int64) *model.FixtureResponse {
	type FixtureParams struct {
		Team   int64 	`url:"team,omitEmpty"`
		Next   int    `url:"next"`
		Status string `url:"status"`
	}

	queryParams := &FixtureParams{
		Team:   teamID,
		Next:   6,
		Status: "NS", // Not Started
	}

	body := &model.FixtureResponse{}

	res, err := s.sling.New().Get("fixtures").QueryStruct(queryParams).Receive(body, nil)
	log.Printf("%v", res.Request.URL)

	if err != nil {
		log.Fatal("Fixture request failed:", err)
	}

	return body
}
