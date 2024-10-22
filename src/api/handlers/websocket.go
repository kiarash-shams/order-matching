package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"order-matching/api/dto"
	"order-matching/matchingo"
	"order-matching/pkg/logging"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var validIntervals = map[string]time.Duration{
    "100ms": 100 * time.Millisecond,
    "500ms": 500 * time.Millisecond,
    "1000ms": time.Second,
}

var upgrader = websocket.Upgrader{
    CheckOrigin: func(r *http.Request) bool {
        return true
    },
}


func GetOrderBookWS(c *gin.Context) {
    conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
    if err != nil {
        fmt.Println("Error upgrading connection:", err)
        return
    }
    defer conn.Close()

    // Get market and interval parameters
    market := c.Query("stream") // For example, BTC-USD
    intervalStr := c.Query("interval") // For example, 1000ms

    // Check if the market exists
    orderBook, exists := markets[market]
    if !exists {
        conn.WriteJSON(gin.H{"error": "Market not found"})
        return
    }

    // Check if the interval is valid
    interval, ok := validIntervals[intervalStr]
    if !ok {
        conn.WriteJSON(gin.H{"error": "Invalid interval, allowed values are: 100ms, 500ms, 1000ms"})
        return
    }

    // Loop to send order book data
    for {
        depth := orderBook.Depth()
        err := conn.WriteJSON(gin.H{
            "market": market,
            "depth":  depth,
        })
        if err != nil {
            fmt.Println("Error sending data:", err)
            break
        }

        // Sleep for the specified interval before sending data again
        time.Sleep(interval)
    }
}

func WebSocketOrderHandler(c *gin.Context,orderBook *matchingo.OrderBook) {
    conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not upgrade connection"})
        return
    }
    defer conn.Close()

    for {
        // دریافت پیام از کلاینت
        _, msg, err := conn.ReadMessage()
        if err != nil {
            break
        }

        // تبدیل پیام به یک ساختار `OrderRequest`
        req := new(dto.OrderRequest)
        if err := json.Unmarshal(msg, req); err != nil {
            conn.WriteMessage(websocket.TextMessage, []byte("Invalid input: "+err.Error()))
            continue
        }

        // پردازش سفارش
        done, err := processOrder(req)
        if err != nil {
            conn.WriteMessage(websocket.TextMessage, []byte("Error: "+err.Error()))
            continue
        }

        // ارسال پاسخ به کلاینت
        response := gin.H{
            "message": "Order processed successfully",
            "done":    done,
            "depth":   orderBook.Depth(),
        }
        responseJSON, _ := json.Marshal(response)
        conn.WriteMessage(websocket.TextMessage, responseJSON)
    }
}

func processOrder(req *dto.OrderRequest) (bool, error) {
    // Find orderBook Market
    orderBook, exists := markets[req.Market]
    if !exists {
        return false, fmt.Errorf("Market not found")
    }

    var order *matchingo.Order

    switch req.OrderKind {
    case "limit":
        if req.Price <= 0 {
            return false, fmt.Errorf("Price must be greater than 0 for limit orders")
        }
        if req.OrderType == "buy" {
            order = matchingo.NewLimitOrder(req.OrderID, matchingo.Buy, matchingo.FromFloat(req.Amount), matchingo.FromFloat(req.Price), "", "")
            logger.Info(logging.Orderbook, logging.NewLimitOrder, "New LimitOrder Buy", nil)
        } else if req.OrderType == "sell" {
            order = matchingo.NewLimitOrder(req.OrderID, matchingo.Sell, matchingo.FromFloat(req.Amount), matchingo.FromFloat(req.Price), "", "")
            logger.Info(logging.Orderbook, logging.NewLimitOrder, "New LimitOrder Sell", nil)
        } else {
            return false, fmt.Errorf("Invalid order type")
        }

    case "market":
        if req.OrderType == "buy" {
            order = matchingo.NewMarketOrder(req.OrderID, matchingo.Buy, matchingo.FromFloat(req.Amount))
            logger.Info(logging.Orderbook, logging.NewMarketOrder, "New MarketOrder Buy", nil)

        } else if req.OrderType == "sell" {
            order = matchingo.NewMarketOrder(req.OrderID, matchingo.Sell, matchingo.FromFloat(req.Amount))
            logger.Info(logging.Orderbook, logging.NewMarketOrder, "New MarketOrder Sell", nil)
        } else {
            return false, fmt.Errorf("Invalid order type")
        }
    default:
        return false, fmt.Errorf("Invalid order kind")
    }

    _, err := orderBook.Process(order)
    if err != nil {
        logger.Error(logging.Orderbook, logging.Process, "Process Failed", nil)
        return false, err
    }

    logger.Info(logging.Orderbook, logging.Process, "Process Success", nil)
	
	

    return true,nil
}




