package test

import (
	"testing"
	"order-matching/matchingo"
	"github.com/google/uuid"
)

type OrderRequest struct {
	OrderID     string    
	OrderType   string
	Market      string
	Price       float64
	Amount      float64
	OrderKind   string
}


func GenerateUniqueOrderID() string {
	return uuid.NewString()
}

func BenchmarkOrderMatching(b *testing.B) {

	// b.N = 10000000 
	initialPrice := 50000.0 // Starting price
	priceIncrement := 10.0   // Price increment for each iteration
	priceDifference := 5.0   // Difference between buy and sell prices

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		// Set the buy price
		buyPrice := initialPrice + float64(i)*priceIncrement
		
		// Set the sell price with a defined difference
		sellPrice := buyPrice + priceDifference

		buyOrder := OrderRequest{
			OrderType: "buy",
			Market:    "BTCUSDT",
			Price:     buyPrice,
			Amount:    0.1,
			OrderKind: "limit",
		}
	
		sellOrder := OrderRequest{
			OrderType: "sell",
			Market:    "BTCUSDT",
			Price:     sellPrice,
			Amount:    0.1,
			OrderKind: "limit",
		}

		// Generate a unique OrderID
		buyOrder.OrderID = GenerateUniqueOrderID()
		sellOrder.OrderID = GenerateUniqueOrderID()

		buy := matchingo.NewLimitOrder(
			buyOrder.OrderID,
			matchingo.Buy,
			matchingo.FromFloat(buyOrder.Amount),
			matchingo.FromFloat(buyOrder.Price),
			"",
			"",
		)

		sell := matchingo.NewLimitOrder(
			sellOrder.OrderID,
			matchingo.Sell,
			matchingo.FromFloat(sellOrder.Amount),
			matchingo.FromFloat(sellOrder.Price),
			"",
			"",
		)
		

	
		b.Logf("Buy Order: %+v", buy)
		b.Logf("Sell Order: %+v", sell)
	}
}