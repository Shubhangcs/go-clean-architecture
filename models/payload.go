package models

// just a simple payload for response

type ErrorPayload struct{
	Message string `json:"message"`
}

type SuccessPayload struct{
	Message string `json:"message"`
	Data any `json:"data"`
}
