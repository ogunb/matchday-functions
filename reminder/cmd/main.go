package main

import (
	"bytes"
	"encoding/json"
	"github.com/ogunb/matchday-functions/reminder"
	"github.com/ogunb/matchday-functions/reminder/model"
	"log"
	"net/http"
)

func main() {
	b, err := json.Marshal(&model.SmsRequest{
		Phone: "123",
		Event: "LOL",
	})

	if err != nil {
		log.Fatal(err)
	}

	req, err := http.NewRequest("POST", "/", bytes.NewBuffer(b))
	if err != nil {
		log.Fatal(err)
	}

	reminder.SmsUser(nil, req)
}
