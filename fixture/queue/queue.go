package queue

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	cloudtasks "cloud.google.com/go/cloudtasks/apiv2"
	tasks "google.golang.org/genproto/googleapis/cloud/tasks/v2"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/ogunb/matchday-functions/fixture/model"
)

func getQueueName() string {
	return fmt.Sprintf("projects/%s/locations/%s/queues/%s", os.Getenv("PROJECT_ID"), os.Getenv("LOCATION"), os.Getenv("QUEUE"))
}

func PurgeQueue() {
	log.Println("Purging queue...")
	ctx := context.Background()
	client, err := cloudtasks.NewClient(ctx)

	if err != nil {
		log.Fatal(err)
	}

	req := &tasks.PurgeQueueRequest{
		Name: getQueueName(),
	}

	_, purgeErr := client.PurgeQueue(ctx, req)

	if purgeErr != nil {
		log.Fatal(purgeErr)
	}

	log.Println("Purged queue successfully.")
}

func CreateTask(match model.Match) {
	log.Println(fmt.Sprintf("Creating task for %s...", match.Event))

	ctx := context.Background()
	client, err := cloudtasks.NewClient(ctx)

	if err != nil {
		log.Fatal(err)
	}

	t, _ := time.Parse(time.RFC3339, match.Timestamp)

	req := &tasks.CreateTaskRequest{
		Parent: getQueueName(),
		Task: &tasks.Task{
			ScheduleTime: &timestamppb.Timestamp{
				Seconds: t.Unix(),
			},
			MessageType: &tasks.Task_HttpRequest{
				HttpRequest: &tasks.HttpRequest{
					HttpMethod: tasks.HttpMethod_POST,
					Url:        os.Getenv("HANDLER_FUNCTION_ENDPOINT"),
					AuthorizationHeader: &tasks.HttpRequest_OidcToken{
						OidcToken: &tasks.OidcToken{
							ServiceAccountEmail: os.Getenv("SERVICE_ACCOUNT_EMAIL"),
						},
					},
				},
			},
		},
	}

	req.Task.GetHttpRequest().Body = []byte(match.Event)

	_, createErr := client.CreateTask(ctx, req)

	if createErr != nil {
		log.Fatal(createErr)
	}

	log.Println(fmt.Sprintf("Created task for %s.", match.Event))
}
