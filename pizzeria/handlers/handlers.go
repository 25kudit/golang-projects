package handlers

import (
	"encoding/json"
	"net/http"
	"pizzeria/kitchen"
	"pizzeria/models"
)

func CreateOrder(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
		return 
	}

	var order models.Order
	json.NewDecoder(r.Body).Decode(&order)
	order.Status = "received"

	kitchen.PlaceOrder(order)

	json.NewEncoder(w).Encode(order)
}