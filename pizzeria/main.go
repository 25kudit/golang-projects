package main

import (
	"log"
	"net/http"
	"pizzeria/handlers"
	"pizzeria/kitchen"
)

func main() {
	kitchen.Init()
	http.HandleFunc("/order", handlers.CreateOrder)

	log.Println("Starting server at port 8080")
	
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Failed to start server")
	}
}