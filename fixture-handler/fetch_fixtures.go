package fixture

import (
	"log"
	"net/http"

	"github.com/ogunb/matchday-functions/fixture/services"
)


func FetchFixtures(w http.ResponseWriter, r *http.Request) {
	teamService := services.NewTeamService()
	firestore := services.NewFirestoreService()
	teams, err := firestore.GetTeamsWithFollowers()

	if err != nil {
		log.Fatal("Getting teams from db failed:", err)
	}


	for _, team := range *teams {
		teamService.CreateTeamEventTasks(team)
	}
}
