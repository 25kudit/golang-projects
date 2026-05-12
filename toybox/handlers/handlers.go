package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"my-api/models"
	"my-api/storage"
	"my-api/worker"
	"net/http"
)

func GetAllMessages(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid Method", http.StatusMethodNotAllowed)
		return 
	}
	reqs, err := storage.GetAllMessages()
	if err != nil {
		log.Println("Database error: ", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return 
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(reqs)
} 

func CreateMessage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
		return
	}

	var input models.Request
	json.NewDecoder(r.Body).Decode(&input)
	w.Header().Set("Content-Type", "application/json")

	response := models.Response {
		Sender: input.Sender,
		Content: input.Content,
		RespFmt: r.URL.Query().Get("resp_fmt"),
	}
	fmt.Printf("Received: %+v\n", input)
	fmt.Printf("Response: %+v\n", response)
	storage.InsertMessage(input.Sender, input.Content)
	worker.InsertQueue(input.Sender)
	json.NewEncoder(w).Encode(response)
}

func Home(w http.ResponseWriter, r *http.Request) {
	msg := models.Message{Sender: "Server", Content: "JSON Hello!"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(msg)
}
