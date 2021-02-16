package services

//
//func createTaskRequest(matchDate string, timeDiff int64) *tasks.CreateTaskRequest{
//	t, _ := time.Parse(time.RFC3339, matchDate)
//	timestamp := t.Unix() - timeDiff
//
//	req := &tasks.CreateTaskRequest{
//		Parent: getQueueName(),
//		Task: &tasks.Task{
//			ScheduleTime: &timestamppb.Timestamp{
//				Seconds: timestamp,
//			},
//			MessageType: &tasks.Task_HttpRequest{
//				HttpRequest: &tasks.HttpRequest{
//					HttpMethod: tasks.HttpMethod_POST,
//					Url:        os.Getenv("HANDLER_FUNCTION_ENDPOINT"),
//					AuthorizationHeader: &tasks.HttpRequest_OidcToken{
//						OidcToken: &tasks.OidcToken{
//							ServiceAccountEmail: os.Getenv("SERVICE_ACCOUNT_EMAIL"),
//						},
//					},
//				},
//			},
//		},
//	}
//
//	return req
//}
//
//func createMatchTodayTask(match model.Match) {
//	req := createTaskRequest(match.Timestamp, THREE_HOURS_IN_UNIX)
//	message := "Üç saat sonra: " + match.Event
//
//	createTask(req, message)
//}
//
//func createMatchNowTask(match model.Match) {
//	req := createTaskRequest(match.Timestamp, FIVE_MINS_IN_UNIX)
//	message := "Beş dakika sonra: " + match.Event
//
//	createTask(req, message)
//}
//
//func createTask(req *tasks.CreateTaskRequest, message string) {
//	ctx := context.Background()
//	client, err := cloudtasks.NewClient(ctx)
//
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	type body struct {
//		Message string `json:"message"`
//	}
//
//	req.Task.GetHttpRequest().Body, _ = json.Marshal(&body{
//		Message: message,
//	})
//
//	_, createErr := client.CreateTask(ctx, req)
//
//	if createErr != nil {
//		log.Fatal(createErr)
//	}
//
//	log.Println(fmt.Sprintf("Created sms task with %s message.", message))
//}
//
//func CreateTask(match model.Match) {
//	log.Println(fmt.Sprintf("Creating task for %s...", match.Event))
//
//	createMatchTodayTask(match)
//	createMatchNowTask(match)
//}