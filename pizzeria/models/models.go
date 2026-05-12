package models

type Order struct{
	Id int 			`json:"id"`
	Item string 	`json:"item"`
	Status string 	`json:"status"`
}