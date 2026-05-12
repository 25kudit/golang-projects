package models

type BinanceResponse struct {
	Symbol string `json:"symbol"`
	Price string  `json:"price"`
}