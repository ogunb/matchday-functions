package main

import reminder "github.com/ogunb/matchday-functions/reminder"

func main() {
	reminder.SmsUser(nil, reminder.PubSubMessage{ Message: []byte("ABC") });
}