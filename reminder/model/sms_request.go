package model

type SmsRequest struct {
	Phone string `json:"phone"`
	Event string `json:"event"`
}