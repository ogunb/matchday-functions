package model

type Status struct {
	Long  string `json:"long"`
	Short string `json:"short"`
}

type Fixture struct {
	Timestamp int64  `json:"timestamp"`
	Status    Status `json:"status"`
}

type Team struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type Teams struct {
	Home Team `json:"home"`
	Away Team `json:"away"`
}

type FixtureResponse struct {
	Response []struct {
		Fixture Fixture `json:"fixture"`
		Teams   Teams   `json:"teams"`
	} `json:"response"`
}
