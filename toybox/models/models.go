package models

type Request struct {
	Sender string `json:"user_id"`
	Content string `json:"req_body"`
}

type Response struct {
	Sender string 
	Content string
	RespFmt string
}

type Message struct {
	Sender string `json:"user_id"`
	Content string `json:"body"`
	RespFmt string	
}