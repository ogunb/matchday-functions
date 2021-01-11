package queue

import (
	"context"
	"fmt"
	"log"

	cloudtasks "cloud.google.com/go/cloudtasks/apiv2"
	tasks "google.golang.org/genproto/googleapis/cloud/tasks/v2"

	"github.com/ogunb/matchday-functions/fixture/config"
	"github.com/ogunb/matchday-functions/fixture/model"
)

func PurgeCloudQueue() {
	log.Println("TODO: PURGE CLOUD QUEUE")
}

func AddToCloudQueue(match model.Match) {
	ctx := context.Background()

	client, err := cloudtasks.NewClient(ctx)

	if err != nil {
		log.Fatal(err)
	}

	configs := config.GetConfig()

	parent := fmt.Sprintf("projects/%s/locations/%s/queues/%s", configs.ProjectID, configs.Location, configs.Queue)

	req := &tasks.CreateTaskRequest{
		Parent: parent,
		Task: &tasks.Task{
			MessageType: &tasks.Task_HttpRequest{
				HttpRequest: &tasks.HttpRequest{
					HttpMethod: tasks.HttpMethod_POST,
					Url:        configs.HandlerFunctionEndpoint,
					AuthorizationHeader: &tasks.HttpRequest_OidcToken{
						OidcToken: &tasks.OidcToken{
							ServiceAccountEmail: configs.ServiceAccountEmail,
						},
					},
				},
			},
		},
	}

	// req.Task.GetAppEngineHttpRequest().Body = []byte(message)

	createdTask, err := client.CreateTask(ctx, req)

	if err != nil {
		log.Fatal(err)
	}

	log.Println(createdTask)
}
