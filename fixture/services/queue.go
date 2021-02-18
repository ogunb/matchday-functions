package services

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/ogunb/matchday-functions/fixture/model"

	cloudtasks "cloud.google.com/go/cloudtasks/apiv2"
	"google.golang.org/genproto/googleapis/cloud/tasks/v2"
)

const (
	threeHoursInUnix = 60 * 60 * 3
	fiveMinsInUnix   = 60 * 5
)

type QueueService struct {
	client *cloudtasks.Client
}

func NewQueueService() *QueueService {
	ctx := context.Background()
	client, err := cloudtasks.NewClient(ctx)

	if err != nil {
		log.Fatal(err)
	}

	return &QueueService{
		client: client,
	}
}

var locationPath = fmt.Sprintf("projects/%s/locations/%s", os.Getenv("PROJECT_ID"), os.Getenv("LOCATION"))

func generateQueueName(team model.Team) string {
	return fmt.Sprintf("%v-%v", team.Name, team.ID)
}

func (s *QueueService) GetLocationPath() string {
	return locationPath
}

func (s *QueueService) GenerateQueuePath(team model.Team) string {
	return fmt.Sprintf("%s/queues/%s", locationPath, generateQueueName(team))
}

func (s *QueueService) GetQueue(queuePath string) (*tasks.Queue, error) {
	ctx := context.Background()
	req := &tasks.GetQueueRequest{Name: queuePath}
	return s.client.GetQueue(ctx, req)
}

func (s *QueueService) DoesQueueExist(queuePath string) bool {
	_, err := s.GetQueue(queuePath)

	if err != nil {
		return false
	}

	return true
}

func (s *QueueService) CreateQueue(queuePath string) {
	ctx := context.Background()
	req := &tasks.CreateQueueRequest{Parent: locationPath, Queue: &tasks.Queue{Name: queuePath}}
	s.client.CreateQueue(ctx, req)
}

func (s *QueueService) PurgeQueue(queuePath string) {
	log.Printf("Purging tasks for %s...\n", queuePath)
	ctx := context.Background()

	req := &tasks.PurgeQueueRequest{
		Name: queuePath,
	}

	_, purgeErr := s.client.PurgeQueue(ctx, req)

	if purgeErr != nil {
		log.Fatal(purgeErr)
	}

	log.Printf("Purged tasks successfully for %s\n", queuePath)

	time.Sleep(2 * time.Second)
}
