package stat

import "time"

type StatPayload struct {
	From string `json:"from"`
	To   string `json:"to"`
	By   string `json:"by" validate:"oneof=day month"`
}

type StatGetResponse struct {
	From time.Time `json:"from"`
	To   time.Time `json:"to"`
	By   string    `json:"by"`
}
