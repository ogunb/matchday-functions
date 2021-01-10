package model

type Match struct {
	Event     string `json:"strEvent"`
	Timestamp string `json:"strTimestamp"`
}

type EventsResponse struct {
	Events []Match `json:"events"`
}