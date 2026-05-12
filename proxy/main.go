package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"proxy/cache"
	"proxy/fetcher"
	"proxy/handlers"
	"syscall"
	"time"
)

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cache := cache.NewPriceCache(ctx)
	fetcher := fetcher.NewBinanceFetcher(nil)
	handler := handlers.NewPriceHandler(fetcher, cache)

	server := &http.Server{Addr: ":8080", Handler: nil}

	http.HandleFunc("/price", handler.GetPrice)

	log.Println("Starting server at 8080")
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("Failed to start server: ", err)
		}
	}()

	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, os.Interrupt, syscall.SIGTERM)
	<-stopChan

	cancel()
	
	shoutDownCtx, shutDownCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutDownCancel()

	if err := server.Shutdown(shoutDownCtx); err != nil {
		log.Fatal("Error while shutting down server: ", err)
	}

	log.Println("Graceful shutdown complete")
}