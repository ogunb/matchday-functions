package services

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ogunb/matchday-functions/fixture/model"
	"google.golang.org/genproto/googleapis/cloud/tasks/v2"
	"google.golang.org/protobuf/types/known/timestamppb"
	"os"

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

func (s *QueueService) CreateTask(queuePath string, teamID int64, fixture model.Fixture, teams model.Teams) {
	ctx := context.Background()

	event := fmt.Sprintf("%s vs. %s", teams.Home.Name, teams.Away.Name)

	log.Println(fmt.Sprintf("Creating task for %s...", event))

	req := createTaskRequest(fixture.Timestamp-FIVE_MINS_IN_UNIX, queuePath)
	type body struct {
		Event  string `json:"event"`
		TeamID int64  `json:"teamId"`
	}

	req.Task.GetHttpRequest().Body, _ = json.Marshal(&body{
		Event:  event,
		TeamID: teamID,
	})

	_, createErr := s.client.CreateTask(ctx, req)

	if createErr != nil {
		log.Fatal(createErr)
	}
}
