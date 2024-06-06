package models

// basic food model with its props
type Food struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	Price  string `json:"price"`
	Ratint int    `json:"rating"`
}


