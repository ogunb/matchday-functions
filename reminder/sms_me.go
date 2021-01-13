package reminder

import (
	"encoding/json"
	"github.com/ogunb/matchday-functions/reminder/model"
	"net/http"
)

// SmsUser smses user?
func SmsUser(res http.ResponseWriter, req *http.Request) {
	var requestBody model.SmsRequest

	if err := json.NewDecoder(req.Body).Decode(&requestBody); err != nil {
		http.Error(res, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}


}
