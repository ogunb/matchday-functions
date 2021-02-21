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
		sling: sling.New().Base(os.Getenv("TEAM_ENDPOINT")),
	}
}

func (s *TeamService) CreateTeamEventTasks(team model.Team) {
	log.Println(team)
}
