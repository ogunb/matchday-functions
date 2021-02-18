package fixture

import (
	"log"
	"net/http"
	"sync"

	"github.com/ogunb/matchday-functions/fixture/model"
	"github.com/ogunb/matchday-functions/fixture/services"
)


func FetchFixtures(w http.ResponseWriter, r *http.Request) {
	firestore := services.NewFirestoreService()
	teams, err := firestore.GetTeamsWithFollowers()

	if err != nil {
		log.Fatal(err)
	}

	var wg sync.WaitGroup
	wg.Add(len(*teams))

	for _, team := range *teams {
		go func(team model.Team) {
			defer wg.Done()
			teamService := services.NewTeamService()
			teamService.CreateTeamEventTasks(team)
		}(team)
	}

	wg.Wait()
}
