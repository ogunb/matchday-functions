package services

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/ogunb/matchday-functions/team-fixture-tasks/apis"
	"github.com/ogunb/matchday-functions/team-fixture-tasks/model"
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

	queuePath := s.GenerateQueuePath(team)
	queueExists := s.queueService.DoesQueueExist(queuePath)

	if !queueExists {
		s.queueService.CreateQueue(queuePath)
		log.Print("Waiting queue creation...")
		time.Sleep(time.Minute)
	} else {
		s.queueService.PurgeQueue(queuePath)
	}


	for _, fixture := range fixture {
		type data struct {
			Event     string `json:"event"`
			TeamID    int64  `json:"teamId"`
			FixtureID int64  `json:"fixtureId"`
			TopicName string `json:"topicName"`
		}

		body := &data{
			Event: fmt.Sprintf("%s vs. %s", fixture.Teams.Home.Name, fixture.Teams.Away.Name),
			TeamID: team.ID,
			FixtureID: fixture.Fixture.ID,
			TopicName: os.Getenv("MATCHDAY_EVENT_TOPIC"),
		}

		scheduleTime := fixture.Fixture.Timestamp - 60 * 5

		s.queueService.CreateTask(queuePath, scheduleTime, body)
	}
}

func (s *TeamService) GenerateQueuePath(team model.Team) string {
	return fmt.Sprintf("%s/queues/%s", locationPath, generateQueueName(team))
}

func generateQueueName(team model.Team) string {
	return fmt.Sprintf("%v-%v", team.Name, team.ID)
}
