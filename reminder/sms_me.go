package reminder

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/ogunb/matchday-functions/reminder/model"
)

// SmsUser smses user?
func SmsUser(res http.ResponseWriter, req *http.Request) {
	accountSid := os.Getenv("TWILIO_SID")
	authToken := os.Getenv("TWILIO_AUTH_TOKEN")
	urlStr := "https://api.twilio.com/2010-04-01/Accounts/" + accountSid + "/Messages.json"

	var requestBody model.SmsRequest

	if err := json.NewDecoder(req.Body).Decode(&requestBody); err != nil {
		http.Error(res, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}

	msgData := url.Values{}
	msgData.Set("From", os.Getenv("TWILIO_NUMBER"))
	msgData.Set("To", os.Getenv("MY_NUMBER"))
	msgData.Set("Body", requestBody.Message)
	msgDataReader := *strings.NewReader(msgData.Encode())

	client := &http.Client{}
	r, _ := http.NewRequest("POST", urlStr, &msgDataReader)
	r.SetBasicAuth(accountSid, authToken)
	r.Header.Add("Accept", "application/json")
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, _ := client.Do(r)
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		var data map[string]interface{}
		decoder := json.NewDecoder(resp.Body)
		err := decoder.Decode(&data)
		if err == nil {
			log.Println(data["sid"])
		}
	} else {
		log.Println(resp.Status)
	}
}
