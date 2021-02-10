package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/ogunb/matchday-functions/reminder"
	"github.com/ogunb/matchday-functions/reminder/model"
)

func main() {
	b, err := json.Marshal(&model.SmsRequest{
		Phone:   os.Getenv("MY_NUMBER"),
		Message: "LOL",
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
