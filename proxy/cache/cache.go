package cache

import (
	"context"
	"log"
	"sync"
	"time"
)

type Price struct {
	price float64
	ttl time.Time
}

type PriceCache struct {
	mu sync.RWMutex
	prices map[string]Price
}

func NewPriceCache(ctx context.Context) *PriceCache {
	pc := &PriceCache{
		prices: make(map[string]Price),
	}
	go pc.startCacheSweeper(5*time.Second, ctx)
	return pc 
}

func (c *PriceCache) startCacheSweeper(interval time.Duration, ctx context.Context) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for {
		select {
		case <- ctx.Done() :
			log.Println("[SWEEPER] Received close signal. Going home!")
			return
		case <- ticker.C :
			c.mu.Lock()
			for key, val := range c.prices {
				if time.Now().After(val.ttl) {
					log.Println("[SWEEPER] Cleaning cache for key: ", key)
					delete(c.prices, key)
				}
			}
			c.mu.Unlock()
		}
	}
}

func (c *PriceCache) Set(s string, p float64) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.prices[s] = Price{price: p, ttl: time.Now().Add(10*time.Second)}
}

func (c *PriceCache) Get(s string) (float64, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	val, ok := c.prices[s]
	if ok && time.Now().After(val.ttl) {
		return -1, false
	}
	return val.price, ok
}