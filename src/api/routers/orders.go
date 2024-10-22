package routers

import (
	"order-matching/api/handlers"
	"order-matching/config"
	"github.com/gin-gonic/gin"
	"order-matching/matchingo"
)



func Order(r *gin.RouterGroup, cfg *config.Config, orderBook *matchingo.OrderBook) {

	r.POST("/create", func(c *gin.Context) {
		handlers.CreateOrder(c)
	})

	r.POST("/orderbook", func(c *gin.Context) {
		handlers.GetOrderBook(c)
	})

	r.POST("/get-order", func(c *gin.Context) {
		handlers.GetOrderById(c)
	})

	r.POST("/cancel-order", func(c *gin.Context) {
		handlers.CancelOrder(c)
	})
}


