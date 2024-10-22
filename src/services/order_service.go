package services

// import (
// 	"sync"
// 	"github.com/gonevo/matchingo"
// )

// type OrderBookService struct {
// 	OrderBook *matchingo.OrderBook
// 	mu        sync.Mutex // برای هماهنگ‌سازی درخواست‌ها
// }

// func NewOrderBookService() *OrderBookService {
// 	return &OrderBookService{
// 		OrderBook: matchingo.NewOrderBook(),
// 	}
// }


// func (obs *OrderBookService) CreateLimitOrder(orderID string, side matchingo.Side, quantity, price float64, tif matchingo.TIF, oco string) (*matchingo.Order, error) {
// 	obs.mu.Lock()
// 	defer obs.mu.Unlock()

// 	// ساخت سفارش جدید
// 	order := matchingo.NewLimitOrder(orderID, side, quantity, price, tif, oco)

// 	// پردازش سفارش
// 	_, err := obs.OrderBook.Process(order)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return order, nil
// }