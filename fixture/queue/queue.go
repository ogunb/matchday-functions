package queue

import (
	"context"
	"fmt"
	"log"
	"os"

	cloudtasks "cloud.google.com/go/cloudtasks/apiv2"
	tasks "google.golang.org/genproto/googleapis/cloud/tasks/v2"

	"github.com/ogunb/matchday-functions/fixture/model"
)

func PurgeQueue() {
	log.Println("TODO: PURGE CLOUD QUEUE")
}

func CreateTask(match model.Match) {
	ctx := context.Background()
	fmt.Println(os.Getenv("PROJECT_ID"))
	client, err := cloudtasks.NewClient(ctx)

	if err != nil {
		log.Fatal(err)
	}

	parent := fmt.Sprintf("projects/%s/locations/%s/queues/%s", os.Getenv("PROJECT_ID"), os.Getenv("LOCATION"), os.Getenv("QUEUE"))

	req := &tasks.CreateTaskRequest{
		Parent: parent,
		Task: &tasks.Task{
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

	// req.Task.GetAppEngineHttpRequest().Body = []byte(message)

	createdTask, err := client.CreateTask(ctx, req)

	if err != nil {
		log.Fatal(err)
	}

	log.Println(createdTask)
}
