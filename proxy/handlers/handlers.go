package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"proxy/cache"
	"proxy/fetcher"
	"time"
)

type PriceHandler struct {
	fetcher *fetcher.BinanceFetcher
	cache *cache.PriceCache
}

func NewPriceHandler(fetcher *fetcher.BinanceFetcher, cache *cache.PriceCache) *PriceHandler {
	return &PriceHandler{fetcher: fetcher, cache: cache}
} 

func (h *PriceHandler) GetPrice(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return 
	}

	symbol := r.URL.Query().Get("symbol")
	if symbol == "" {
		http.Error(w, "Missing symbol parameter in request", http.StatusBadRequest)
		return 
	}

	price, ok := h.cache.Get(symbol)
	w.Header().Set("Content-Type", "application/json")

	if ok {
		log.Println("[HIT] Price found in cache ", symbol, ":", price)
		json.NewEncoder(w).Encode(price)
		return
	}

	log.Println("[MISS] Price not found in cache for ", symbol)

	ctx, cancel := context.WithTimeout(r.Context(), 3*time.Second)
	defer cancel()

	price, err := h.fetcher.FetchPrice(ctx, symbol)
	if err != nil {
		log.Println("[ERROR] Failed to fetch price from api ", err)
		http.Error(w, "Failed to fetch price", http.StatusInternalServerError)
		return 
	}
	h.cache.Set(symbol,price)
	json.NewEncoder(w).Encode(price)

}