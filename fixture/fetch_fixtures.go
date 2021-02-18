package fixture

import (
	"net/http"

	"github.com/ogunb/matchday-functions/fixture/model"
	"github.com/ogunb/matchday-functions/fixture/services"
)


func FetchFixtures(w http.ResponseWriter, r *http.Request) {
	teamService := services.NewTeamService()
	team := &model.Team{ID: 549, Name: "Besiktas"}
	teamService.CreateTeamEventTasks(*team)
}
