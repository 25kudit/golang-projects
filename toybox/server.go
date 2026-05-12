package main

import (
	// "encoding/json"
	// "fmt"
	"context"
	"log"
	"my-api/handlers"
	"my-api/storage"
	"my-api/worker"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	storage.InitDB()
	worker.InitTaskQueue()

	server := &http.Server{Addr: ":8080", Handler: nil}

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/create", handlers.CreateMessage)
	http.HandleFunc("/messages", handlers.GetAllMessages)

	go func() {
		log.Println("Starting server at port 8080")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("Failed with error: ", err)
		}
	}()

	stopChan := make (chan os.Signal, 1)
	signal.Notify(stopChan, os.Interrupt, syscall.SIGTERM)
	<-stopChan
	log.Println("Received stop signal, shutting down")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Error while server shutdown", err)
	}

	log.Println("HTTP shutdown success")

	worker.Close()

	log.Println("Graceful shutdown complete")

}
