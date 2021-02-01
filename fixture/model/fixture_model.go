package model

type Match struct {
	Event     string `json:"strEvent"`
	Timestamp string `json:"strTimestamp"`
	Status    string `json:"strStatus"`
}

type EventsResponse struct {
	Events []Match `json:"events"`
}
