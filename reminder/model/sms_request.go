package model

type SmsRequest struct {
	Phone   string `json:"phone"`
	Message string `json:"message"`
}
