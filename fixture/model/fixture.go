package model

type Fixture struct {
	Timestamp int `json:"timestamp"`
	Status    struct {
		Long string `json:"long"`
	} `json:"status"`
}

type Team struct {
	ID int `json:"id"`
	Name string `json:"name"`
}

type FixtureResponse struct {
	Response []struct {
		Fixture Fixture `json:"fixture"`
		Teams   struct {
			Home Team `json:"home"`
			Away Team `json:"away"`
		} `json:"teams"`
	} `json:"response"`
}
