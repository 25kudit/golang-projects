package main

import (
	"context"
	"crawler/engine"
	"crawler/models"
	"crawler/storage"
	"log"
	"sync"
)


func worker(
	id int,
	wg *sync.WaitGroup,
	c *engine.Crawler,
	jobs <-chan string,
	results chan<- models.PageResult,
) {
	defer wg.Done()
	for url := range jobs {
		log.Printf("[WORKER %d] Processing url: %s\n", id, url)
		ctx := context.Background()
		res := c.Fetch(ctx, url)
		results <- res
	}
	log.Printf("[WORKER %d] Finished and retiring.\n", id)
}

func main(){
	urls := []string{
		"https://golang.org", 
		"https://google.com", 
		"https://github.com", 
		"https://pkg.go.dev", 
		"https://reddit.com",
	}

	crawler := engine.NewCrawler(nil)
	mStore := storage.NewMemoryStore()

	jobs := make(chan string, len(urls))
	results := make(chan models.PageResult, 3)
	saveDone := make(chan bool)

	var wg sync.WaitGroup

	go func() {
		for res := range results {
			mStore.Save(res)
			log.Printf("[SAVER] Saved result for %s (Length: %d)\n", res.URL, res.BodyLength)
		}
		saveDone <- true
	}()

	for i := range 3 {
		wg.Add(1)
		go worker(i, &wg, crawler, jobs, results)
	}

	for _, url := range urls {
		jobs <- url
	}
	close(jobs)

	wg.Wait()
	close(results)
	<-saveDone
	log.Println("Crawler finished successfully")

}