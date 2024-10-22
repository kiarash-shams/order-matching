package handlers

import (
	"fmt"
	"net/http"
	"order-matching/api/dto"
	"github.com/gin-gonic/gin"
	"order-matching/matchingo"
	"order-matching/pkg/logging"
)

// var logger = logging.NewLogger(config.GetConfig())

var markets = map[string]*matchingo.OrderBook{
	"BTCIRT": matchingo.NewOrderBook(),
	"ETHIRT": matchingo.NewOrderBook(),
	"USDTIRT": matchingo.NewOrderBook(),
	"BTCUSDT": matchingo.NewOrderBook(),
	"ETHUSDT": matchingo.NewOrderBook(),
}


// CreateOrder godoc
// @Summary Create an order
// @Description Create a new order of type limit or market
// @Tags Orders
// @Accept  json
// @Produce  json
// @Param Request body dto.OrderRequest true "OrderRequest"
// @Success 200 {object} gin.H "Order processed successfully"
// @Failure 400 {object} gin.H "Invalid input or processing error"
// @Router /v1/orders/create [post]
func CreateOrder(c *gin.Context,) {
	req := new(dto.OrderRequest)
	
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: " + err.Error()})
		return
	}

	// Find orderBook Market
	orderBook, exists := markets[req.Market]
	if !exists {
		logger.Error(logging.Orderbook, logging.MarketNotFound, "Market not found", nil)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Market not found"})
		return
	}
	
	var order *matchingo.Order

	switch req.OrderKind {
	case "limit":
		
		if req.Price <= 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Price must be greater than 0 for limit orders"})
			return
		}
		if req.OrderType == "buy" {
			order = matchingo.NewLimitOrder(req.OrderID, matchingo.Buy, matchingo.FromFloat(req.Amount), matchingo.FromFloat(req.Price), "", "")
			logger.Info(logging.Orderbook, logging.NewLimitOrder, "New LimitOrder Buy", nil)
		} else if req.OrderType == "sell" {
			order = matchingo.NewLimitOrder(req.OrderID, matchingo.Sell, matchingo.FromFloat(req.Amount), matchingo.FromFloat(req.Price), "", "")
			logger.Info(logging.Orderbook, logging.NewLimitOrder, "New LimitOrder Sell", nil)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order type"})
			return
		}

	case "market":
		if req.OrderType == "buy" {
			order = matchingo.NewMarketOrder(req.OrderID, matchingo.Buy, matchingo.FromFloat(req.Amount))
			logger.Info(logging.Orderbook, logging.NewMarketOrder, "New MarketOrder Buy", nil)
			
		} else if req.OrderType == "sell" {
			order = matchingo.NewMarketOrder(req.OrderID, matchingo.Sell, matchingo.FromFloat(req.Amount))	
			logger.Info(logging.Orderbook, logging.NewMarketOrder, "New MarketOrder Sell", nil)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order type"})
			return
		}
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order kind"})
		return
	}


	done, err := orderBook.Process(order)
	if err != nil {
		logger.Error(logging.Orderbook, logging.Process, "Process Orderbook Failed", nil)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return}
	
	logger.Info(logging.Orderbook, logging.Process, "Process Orderbook Success", nil)

	c.JSON(http.StatusOK, gin.H{
		"message": "Order processed successfully",
		"done":    done,
		"depth": orderBook.Depth(),				
	})
}

// CancelOrder godoc
// @Summary Cancel an order
// @Description Cancel an order by its ID and market
// @Tags Orders
// @Accept  json
// @Produce  json
// @Param Request body dto.CancelOrderRequest true "CancelOrderRequest"
// @Success 200 {object} gin.H "Order canceled successfully"
// @Failure 400 {object} gin.H "Invalid input"
// @Failure 404 {object} gin.H "Market or Order not found"
// @Router /v1/orders/cancel-order [post]
func CancelOrder(c *gin.Context) {
    req := new(dto.CancelOrderRequest)

    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: " + err.Error()})
        return
    }

    orderBook, exists := markets[req.Market]
    if !exists {
        c.JSON(http.StatusNotFound, gin.H{"error": "Market not found"})
        return
    }

	//cancel Order
    canceled := orderBook.CancelOrder(req.OrderID)
    if canceled == nil{
		logger.Error(logging.Orderbook, logging.OrderCancelled, "OrderCancel", nil)
        c.JSON(http.StatusNotFound, gin.H{"error": "Order not found or already canceled"})
        return
    }

	logger.Info(logging.Orderbook, logging.OrderCancelled, "OrderCancel", nil)


    c.JSON(http.StatusOK, gin.H{
        "message": "Order canceled successfully",
    })
}

// GetOrderById godoc
// @Summary Get an order by ID
// @Description Retrieve order details by its ID and market
// @Tags Orders
// @Accept  json
// @Produce  json
// @Param Request body dto.GetOrderRequest true "GetOrderRequest"
// @Success 200 {object} gin.H "Order Find"
// @Failure 400 {object} gin.H "Invalid input"
// @Failure 404 {object} gin.H "Market or Order not found"
// @Router /v1/orders/get-order [post]
func GetOrderById(c *gin.Context) {
	req := new(dto.GetOrderRequest)
	
	if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: " + err.Error()})
        return
    }

    // Find Market exits
    orderBook, exists := markets[req.Market]
    if !exists {
        c.JSON(http.StatusNotFound, gin.H{"error": "Market not found"})
        return
    }
	 // Find OrderById 
	order := orderBook.GetOrder(req.OrderID)
	 if order == nil {
		 c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		 return
	 }
 
	logger.Info(logging.Orderbook, logging.GetOrderId, "GetOrderId", nil)
	c.JSON(http.StatusOK, gin.H{
		"message": "ّOrder Find",
		"order": order.String(),
	})
}

// GetOrderBook godoc
// @Summary Get order book depth
// @Description Retrieve the depth of the order book for a specific market
// @Tags Orders
// @Accept  json
// @Produce  json
// @Param Request body dto.GetOrderBookRequest true "GetOrderBookRequest"
// @Success 200 {object} gin.H "OrderBook depth retrieved"
// @Failure 400 {object} gin.H "Invalid input"
// @Failure 404 {object} gin.H "Market not found"
// @Router /v1/orders/orderbook [post]
func GetOrderBook(c *gin.Context) {
	req := new(dto.GetOrderBookRequest)
	
	// Validate the request body to ensure "market" is provided
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: " + err.Error()})
		return
	}

	// Check if the market exists in the orderBook
	orderBook, exists := markets[req.Market]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Market not found"})
		return
	}

	logger.Info(logging.Orderbook, logging.ListOrderBook, "ListOrderBook", nil)
	// Return the depth of the orderBook for the given market
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("OrderBook %s live", req.Market),
		"depth":   orderBook.Depth(),
	})
}