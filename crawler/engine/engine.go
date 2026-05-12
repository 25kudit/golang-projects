package engine

import (
	"context"
	"crawler/models"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Crawler struct {
	client *http.Client
}

func NewCrawler(c *http.Client) *Crawler {
	if c == nil {
		c = &http.Client{}
	}
	return &Crawler{client: c}
}

func (c *Crawler) Fetch(ctx context.Context, targetURL string) models.PageResult {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, targetURL, nil)
	if err != nil {
		log.Println("[ENGINE] Error while creating http request: ", err)
		return models.PageResult{URL: targetURL, Error: err.Error()}
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36")

	resp,err := c.client.Do(req)

	if err != nil {
		log.Println("[ENGINE] Fetch failed: ", err)
		return models.PageResult{URL: targetURL, Error: err.Error()}
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Println("[ENGINE] Get request unsuccessfull. Status code: ", resp.StatusCode)
		return models.PageResult{URL: targetURL, Error: fmt.Errorf("received non-200 status code").Error()}
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("[ENGINE] Failed while reading response body: ", err)
		return models.PageResult{URL: targetURL, Error: err.Error()}
	}

	return models.PageResult{URL: targetURL, Title: "N/A", BodyLength: len(body), Error: ""}
}