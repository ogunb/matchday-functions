package services

import (
	"context"
	"fmt"
	"github.com/ogunb/matchday-functions/fixture/model"
	"log"
	"os"
	"time"

	cloudtasks "cloud.google.com/go/cloudtasks/apiv2"
	"google.golang.org/genproto/googleapis/cloud/tasks/v2"
)

const (
	THREE_HOURS_IN_UNIX = 60 * 60 * 3
	FIVE_MINS_IN_UNIX   = 60 * 5
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
	return fmt.Sprintf("%v-%s", team.ID, team.Name)
}

func (s *QueueService) GetLocationPath() string {
	return locationPath
}

func (s *QueueService) GenerateQueuePath(team model.Team) string {
	return fmt.Sprintf("%s/queues/%s", locationPath, generateQueueName(team))
}

func (s *QueueService) CreateQueue(queuePath string) {
	ctx := context.Background()
	req := &tasks.CreateQueueRequest{Parent: locationPath, Queue: &tasks.Queue{Name: queuePath}}
	_, err := s.client.CreateQueue(ctx, req)
	if err != nil {
		log.Println(err.Error())
	}
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
