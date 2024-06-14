package models

//Food order model
type Orders struct{
	OrderId string `json:"orderid"`
	OrderName string `json:"ordername"`
	OrderAddress string `json:"orderaddress"`
	OrderPrice string `json:"orderprice"`
}