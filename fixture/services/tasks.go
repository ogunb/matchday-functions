package services

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"google.golang.org/genproto/googleapis/cloud/tasks/v2"
	"google.golang.org/protobuf/types/known/timestamppb"

	"log"
)

func createTaskRequest(timestamp int64, queuePath string) *tasks.CreateTaskRequest {
	req := &tasks.CreateTaskRequest{
		Parent: queuePath,
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

func (s *QueueService) CreateTask(queuePath string, timestamp int64, data interface{}) {
	ctx := context.Background()

	log.Println(fmt.Sprintf("Creating task for %s...", queuePath))

	req := createTaskRequest(timestamp, queuePath)

	req.Task.GetHttpRequest().Body, _ = json.Marshal(data)

	_, createErr := s.client.CreateTask(ctx, req)

	if createErr != nil {
		log.Fatal("Creating task failed:", createErr)
	}
}
