package apis

import (
	"github.com/dghubble/sling"
	"github.com/ogunb/matchday-functions/fixture/model"
	"log"
	"os"
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

func (s *SportsService) FetchNotStartedMatches(teamID string) *model.FixtureResponse {
	type FixtureParams struct {
		Team   string `url:"team,omitEmpty"`
		Next   int    `url:"next"`
		Status string `url:"status"`
	}

	queryParams := &FixtureParams{
		Team:   teamID,
		Next:   10,
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
