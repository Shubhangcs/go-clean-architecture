package models

// just a simple payload for response

type SimplePayload struct {
	Message string `json:"message"`
	Status  int64  `json:"status"`
}
