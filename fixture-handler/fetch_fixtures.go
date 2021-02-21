package fixture

import (
	"log"
	"net/http"
	"sync"

	"github.com/ogunb/matchday-functions/fixture/model"
	"github.com/ogunb/matchday-functions/fixture/services"
)


func FetchFixtures(w http.ResponseWriter, r *http.Request) {
	teamService := services.NewTeamService()
	firestore := services.NewFirestoreService()
	teams, err := firestore.GetTeamsWithFollowers()

	if err != nil {
		log.Fatal("Getting teams from db failed:", err)
	}

	var wg sync.WaitGroup
	wg.Add(len(*teams))

	for _, team := range *teams {
		go func(team model.Team) {			teamService.CreateTeamEventTasks(team)
			defer wg.Done()
			teamService.CreateTeamEventTasks(team)
		}(team)
	}

	wg.Wait()
}
