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

func getQueuePath(queueName string) string {
	return fmt.Sprintf("%s/queues/%s", locationPath, queueName)
}

func (s *QueueService) GenerateQueueName(team model.Team) string {
	return fmt.Sprintf("%v-%s", team.ID, team.Name)
}

func (s *QueueService) CreateQueue(team model.Team) {
	ctx := context.Background()
	queueName := s.GenerateQueueName(team)
	req := &tasks.CreateQueueRequest{Parent: locationPath, Queue: &tasks.Queue{Name: getQueuePath(queueName)}}
	_, err := s.client.CreateQueue(ctx, req)
	if err != nil {
		log.Println(err.Error())
	}
}

func (s *QueueService) PurgeQueue(queueName string) {
	log.Printf("Purging tasks for %s...\n", queueName)
	ctx := context.Background()

	req := &tasks.PurgeQueueRequest{
		Name: getQueuePath(queueName),
	}

	_, purgeErr := s.client.PurgeQueue(ctx, req)

	if purgeErr != nil {
		log.Fatal(purgeErr)
	}

	log.Printf("Purged tasks successfully for %s\n", queueName)

	time.Sleep(2 * time.Second)
}
