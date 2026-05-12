package worker

import (
	"log"
	"sync"
	"time"
)

var taskQueue chan string
var wg sync.WaitGroup

func InitTaskQueue() {
	taskQueue = make(chan string, 3)
	wg.Add(1)
	go processQueue()
}

func processQueue() {
	defer wg.Done()
	for msg := range taskQueue {
		HeavyTask(msg)
	}
	log.Println("task queue is now closed")
}

func InsertQueue(msg string) {
	taskQueue <- msg
}

func HeavyTask(msg string) {
	log.Println("Starting heavy task")
	log.Println("Processing msg: ", msg)
	time.Sleep(10 * time.Second)
	log.Println("Finished heavy task")
}

func Close() {
	close(taskQueue)
	wg.Wait()
}
