package queue

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	cloudtasks "cloud.google.com/go/cloudtasks/apiv2"
	"google.golang.org/genproto/googleapis/cloud/tasks/v2"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/ogunb/matchday-functions/fixture/model"
)

const (
	THREE_HOURS_IN_UNIX = 60 * 60 * 3
	FIVE_MINS_IN_UNIX = 60 * 5
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

func createTaskRequest(matchDate string, timeDiff int64) *tasks.CreateTaskRequest{
	t, _ := time.Parse(time.RFC3339, matchDate)
	timestamp := t.Unix() - timeDiff

	req := &tasks.CreateTaskRequest{
		Parent: getQueueName(),
		Task: &tasks.Task{
			ScheduleTime: &timestamppb.Timestamp{
				Seconds: timestamp,
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

	return req
}

func createMatchTodayTask(match model.Match) {
	req := createTaskRequest(match.Timestamp, THREE_HOURS_IN_UNIX)
	message := "Üç saat sonra: " + match.Event

	createTask(req, message)
}

func createMatchNowTask(match model.Match) {
	req := createTaskRequest(match.Timestamp, FIVE_MINS_IN_UNIX)
	message := "Beş dakika sonra: " + match.Event

	createTask(req, message)
}

func createTask(req *tasks.CreateTaskRequest, message string) {
	ctx := context.Background()
	client, err := cloudtasks.NewClient(ctx)

	if err != nil {
		log.Fatal(err)
	}

	type body struct {
		Message string `json:"message"`
	}

	req.Task.GetHttpRequest().Body, _ = json.Marshal(&body{
		Message: message,
	})

	_, createErr := client.CreateTask(ctx, req)

	if createErr != nil {
		log.Fatal(createErr)
	}

	log.Println(fmt.Sprintf("Created sms task with %s message.", message))
}

func CreateTask(match model.Match) {
	log.Println(fmt.Sprintf("Creating task for %s...", match.Event))

	createMatchTodayTask(match)
	createMatchNowTask(match)
}
