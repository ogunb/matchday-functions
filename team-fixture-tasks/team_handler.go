package team

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/ogunb/matchday-functions/team-fixture-tasks/model"
	"github.com/ogunb/matchday-functions/team-fixture-tasks/services"
)


func CreateTasks(w http.ResponseWriter, r *http.Request) {
	teamService := services.NewTeamService()

	log.Println(r.Body)
	var team model.Team

	if err := json.NewDecoder(r.Body).Decode(&team); err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		log.Fatal("Error parsing request body ", r.Body)
	}

	teamService.CreateTeamEventTasks(team)
}
