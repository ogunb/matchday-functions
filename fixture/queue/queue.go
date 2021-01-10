package queue

import (
	"context"
	"fmt"
	"log"

	"github.com/ogunb/matchday-functions/fixture/config"
	cloudtasks "google.golang.org/api/cloudtasks/v2"
)

func PurgeCloudQueue() {
	log.Println("TODO: PURGE CLOUD QUEUE")
}

func AddToCloudQueue() {
	ctx := context.Background()
	log.Println(ctx)

	cloudtasksService, err := cloudtasks.NewService(ctx)

	if err != nil {
		log.Fatal(err)
	}

	configs := config.GetConfig()

	name := fmt.Sprintf("projects/%s/locations/%s/queues/%s", configs.ProjectID, configs.Location, configs.Queue)

	req := &cloudtasks.CreateTaskRequest{
		Task: &cloudtasks.Task{
			HttpRequest: &cloudtasks.HttpRequest{
				HttpMethod: "POST",
				Url: configs.HandlerFunctionEndpoint,
				OidcToken: &cloudtasks.OidcToken{
					ServiceAccountEmail: configs.ServiceAccountEmail,
				},
			},
		},
	}

	fmt.Println(req.Task.HttpRequest.OidcToken.ServiceAccountEmail)

	// req.Task.GetAppEngineHttpRequest().Body = []byte(message)

	createdTask := cloudtasksService.Projects.Locations.Queues.Tasks.Create(name, req)

	log.Println(createdTask)
}
