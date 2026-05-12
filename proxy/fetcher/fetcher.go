package fetcher

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"proxy/models"
	"strconv"
)

type BinanceFetcher struct {
	client *http.Client
}

func NewBinanceFetcher(client *http.Client) *BinanceFetcher {
	if client == nil {
		client = &http.Client{}
	}
	return &BinanceFetcher{client: client}
}

func (fetcher *BinanceFetcher) FetchPrice(ctx context.Context, symbol string) (float64, error) {
	reqUrl := "https://api.binance.com/api/v3/ticker/price?symbol=" + symbol
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, reqUrl, nil); 
	if err != nil {
		return -1, err
	}

	httpResp, err := fetcher.client.Do(req)
	if err != nil {
		return -1, err
	}
	defer httpResp.Body.Close()

	if httpResp.StatusCode != http.StatusOK {
		return -1, fmt.Errorf("api return code: %d", httpResp.StatusCode)
	}

	var resp models.BinanceResponse
	if err := json.NewDecoder(httpResp.Body).Decode(&resp); err != nil {
		return -1, err
	}
	
	return strconv.ParseFloat(resp.Price, 64)
}