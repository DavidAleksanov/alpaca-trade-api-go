package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/alpacahq/alpaca-trade-api-go/v2/alpaca"
	talib "github.com/markcheno/go-talib"
)

const (
	SYMBOL   = "AAPL"
	INTERVAL = "15Min"
	LIMIT    = 200
	QNTY     = 35
)

var (
	apiKey    = os.Getenv("ALPACA_API_KEY")
	apiSecret = os.Getenv("ALPACA_API_SECRET")
	baseURL   = "https://paper-api.alpaca.markets"
	client    = alpaca.NewClient(alpaca.ClientOpts{
		APIKey:    apiKey,
		APISecret: apiSecret,
		BaseURL:   baseURL,
	})
)

func getData() ([]float64, error) {
	end := time.Now()
	start := end.Add(-24 * time.Hour)
	bars, err := client.GetBars(SYMBOL, alpaca.ListBarParams{
		TimeFrame: INTERVAL,
		StartDt:   start,
		EndDt:     end,
		Limit:     LIMIT,
	})
	if err != nil {
		return nil, err
	}

	var closes []float64
	for _, bar := range bars {
		closePrice, _ := strconv.ParseFloat(bar.Close, 64)
		closes = append(closes, closePrice)
	}
	return closes, nil
}

func placeBuyOrder() {
	_, err := client.PlaceOrder(alpaca.PlaceOrderRequest{
		AssetKey:    &SYMBOL,
		Qty:         QNTY,
		Side:        alpaca.Buy,
		Type:        alpaca.Market,
		TimeInForce: alpaca.GTC,
	})
	if err != nil {
		log.Printf("Error placing buy order: %s\n", err)
	} else {
		log.Println("Buy order placed")
	}
}

func placeSellOrder() {
	_, err := client.PlaceOrder(alpaca.PlaceOrderRequest{
		AssetKey:    &SYMBOL,
		Qty:         QNTY,
		Side:        alpaca.Sell,
		Type:        alpaca.Market,
		TimeInForce: alpaca.GTC,
	})
	if err != nil {
		log.Printf("Error placing sell order: %s\n", err)
	} else {
		log.Println("Sell order placed")
	}
}

func main() {
	buy := false
	sell := true
	log.Println("Script running...")
	for {
		closingData, err := getData()
		if err != nil {
			log.Fatalf("Error getting data: %s\n", err)
		}

		rsi := talib.Rsi(closingData, 7)
		currentRsi := rsi[len(rsi)-1]
		fmt.Printf("Current RSI: %.2f\n", currentRsi)

		if currentRsi <= 30 && !buy {
			placeBuyOrder()
			buy = true
			sell = false
		}

		if currentRsi >= 70 && !sell {
			placeSellOrder()
			buy = false
			sell = true
		}

		time.Sleep(60 * time.Second)
	}
}
