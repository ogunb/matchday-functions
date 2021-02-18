package services

import (
	"log"
	"time"

	"github.com/ogunb/matchday-functions/fixture/apis"
	"github.com/ogunb/matchday-functions/fixture/model"
)


type TeamService struct {
	sportsApi *apis.SportsService
	queueService *QueueService
}

func NewTeamService() *TeamService {
	return &TeamService{
		sportsApi: apis.NewSportsService(),
		queueService: NewQueueService(),
	}
}

func (s *TeamService) CreateTeamEventTasks(team model.Team) {
	fixture := s.sportsApi.FetchNotStartedMatches(team.ID).Response

	if len(fixture) == 0 {
		log.Fatal("No event was found.")
	}

	queuePath := s.queueService.GenerateQueuePath(team)
	queueExists := s.queueService.DoesQueueExist(queuePath)

	if !queueExists {
		s.queueService.CreateQueue(queuePath)
		log.Print("Waiting queue creation...")
		time.Sleep(time.Minute)
	} else {
		s.queueService.PurgeQueue(queuePath)
	}


	for _, fixture := range fixture {
		s.queueService.CreateTask(queuePath, team.ID, fixture.Fixture, fixture.Teams)
	}
}