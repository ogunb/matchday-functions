package services

import (
	"log"
	"os"

	"github.com/dghubble/sling"
	"github.com/ogunb/matchday-functions/fixture/model"
)


type TeamService struct {
	sling *sling.Sling
}

func NewTeamService() *TeamService {
	return &TeamService{
		sling: sling.New().Base(os.Getenv("PROJECT_URL")),
	}
}

func (s *TeamService) CreateTeamEventTasks(team model.Team) {
	res, err := s.sling.New().Post("team-fixture-handler").BodyJSON(team).Receive(nil, nil)
	log.Printf("%v", res.Request.URL)
	log.Println(res)

	if err != nil {
		log.Fatal("Fixture request failed:", err)
	}
}
