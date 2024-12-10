package api

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
	"order-matching/api/handlers"
	"order-matching/api/middleware"
	"order-matching/api/routers"
	"order-matching/api/validation"
	"order-matching/config"
	"order-matching/docs"
	"order-matching/matchingo"
	"order-matching/pkg/logging"
	"order-matching/pkg/metrics"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/websocket"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"golang.org/x/exp/rand"
)

var logger = logging.NewLogger(config.GetConfig())

type OrderRequest struct {
	Amount   float64 `json:"amount"`
	Market   string  `json:"market"`
	OrderID  string  `json:"order_id"`
	OrderKind string `json:"order_kind"`
	OrderType string `json:"order_type"`
	Price    float64 `json:"price"`
}

var orderBook *matchingo.OrderBook

func InitServer(cfg *config.Config) {


	gin.SetMode(cfg.Server.RunMode)
	r := gin.New()

	RegisterValidators()
	RegisterPrometheus()

	orderBook = matchingo.NewOrderBook()
	matchingo.SetDecimalFraction(5) 
	logger.Info(logging.Orderbook, logging.Startup, "Run OrderBook", nil)
	

	r.Use(middleware.DefaultStructuredLogger(cfg))
	r.Use(middleware.Cors(cfg))
	r.Use(middleware.Prometheus())
	r.Use(gin.Logger(), gin.CustomRecovery(middleware.ErrorHandler))

	RegisterRoutes(r, cfg)	
	RegisterWebSocket(r, cfg)
	logger.Info(logging.WebSocket, logging.Startup, "Run WebSockets ", nil)

	RegisterSwagger(r, cfg)

	// go binance()


	err := r.Run(fmt.Sprintf(":%s", cfg.Server.InternalPort))
	
	if err != nil {
		logger.Error(logging.General, logging.Startup, err.Error(), nil)
	}
}


func sendOrder(order OrderRequest) error {
	// url := "http://localhost:5005/api/v1/orders/create"
	url := "http://localhost:5000/api/v1/orders/create"
	
	orderJSON, err := json.Marshal(order)
	if err != nil {
		return fmt.Errorf("error marshalling order data: %v", err)
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(orderJSON))
	if err != nil {
		return fmt.Errorf("error sending order request: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error reading response body: %v", err)
	}

	// fmt.Printf("Order Response Status: %v\n", resp.Status)
	// fmt.Printf("Order Response Body: %s\n", string(body))

    var responseMap map[string]interface{}
    if err := json.Unmarshal(body, &responseMap); err != nil {
        return fmt.Errorf("error unmarshalling response body: %v", err)
    }

    if done, exists := responseMap["done"].(map[string]interface{}); exists {
        if trades, found := done["trades"].([]interface{}); found {
            fmt.Println("Trades:")
            for _, trade := range trades {
                tradeData := trade.(map[string]interface{})
                fmt.Printf("Trade OrderID: %s, Price: %s, Quantity: %s\n",
                    tradeData["orderID"], tradeData["price"], tradeData["quantity"])
					logger.Info(logging.Orderbook, logging.OrderExecuted, "OrderExecuted", nil)
					metrics.TradeCompleted.WithLabelValues("TradeExecuted", "Success").Inc()
            }
        } else {
            fmt.Println("No trades found.")
        }
    } else {
        fmt.Println("No 'done' or 'trades' section in response.")
    }


	return nil
}

func binance(){
	fmt.Println("Binance WebSocket Run")
	logger.Info(logging.WebSocket, logging.Startup, "Binance WebSocket Run", nil)

	dialer := websocket.Dialer{
        TLSClientConfig: &tls.Config{
            InsecureSkipVerify: true,
        },
    }
	url := "wss://stream.testnet.binance.vision:9443/ws/btcusdt@depth@100ms"
	
	conn, _, err := dialer.Dial(url, nil)
	if err != nil {
		log.Printf("Error connecting to WebSocket: %v", err)
		return 
	}
	defer conn.Close()

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Printf("Error reading message: %v", err)
			return
		}
		// fmt.Println(string(message))
		
		// time.Sleep(100 * time.Millisecond)
		var data map[string]interface{}
		if err := json.Unmarshal(message, &data); err != nil {
			log.Printf("Error unmarshalling message: %v\n", err)
			continue
		}

		bids := data["b"].([]interface{}) // List of buy bids
		asks := data["a"].([]interface{}) // List of sell asks

		for _, bid := range bids {
			bidData := bid.([]interface{})
			bidPrice, _ := strconv.ParseFloat(bidData[0].(string), 64)
			bidAmount, _ := strconv.ParseFloat(bidData[1].(string), 64)

			orderBuy := OrderRequest{
				Amount:   bidAmount,
				Market:   "BTCUSDT",
				OrderID:  generateOrderID(),
				OrderKind: "limit",
				OrderType: "buy",
				Price:    bidPrice,
			}

			if err := sendOrder(orderBuy); err != nil {
				log.Printf("Error sending buy order: %v", err)
				// logger.Error(logging.Orderbook, logging.NewLimitOrder, "Submit Buy Order to Orderbook", nil)
				// metrics.NewOrder.WithLabelValues("CreateBuyOrder", "Failed ").Inc()
			}else {
				log.Printf("Buy order sent: %+v\n", orderBuy)
				// logger.Info(logging.Orderbook, logging.NewLimitOrder, "Submit Buy Order to Orderbook", nil)
				// metrics.NewOrder.WithLabelValues( "CreateBuyOrder", "Success").Inc()
			}
		}

		for _, ask := range asks {
			askData := ask.([]interface{})
			askPrice, _ := strconv.ParseFloat(askData[0].(string), 64)
			askAmount, _ := strconv.ParseFloat(askData[1].(string), 64)

			orderSell := OrderRequest{
				Amount:   askAmount,
				Market:   "BTCUSDT",
				OrderID:  generateOrderID(),
				OrderKind: "limit",
				OrderType: "sell",
				Price:    askPrice,
			}

			if err := sendOrder(orderSell); err != nil {
				log.Printf("Error sending sell order: %v", err)
				// logger.Error(logging.Orderbook, logging.NewLimitOrder, "Submit Sell Order to Orderbook", nil)
				// metrics.NewOrder.WithLabelValues("Sell", "Create", "Failed ").Inc()
				// metrics.NewOrder.WithLabelValues(reflect.TypeOf(err).String(), "CreateSellOrder", "Failed ").Inc()
			} else {
				log.Printf("Sell order sent: %+v\n", orderSell)
				// logger.Info(logging.Orderbook, logging.NewLimitOrder, "Submit Sell Order to Orderbook", nil)
				// metrics.NewOrder.WithLabelValues("Sell", "Create", "Success").Inc()
				// metrics.NewOrder.WithLabelValues(reflect.TypeOf(orderSell).String(), "CreateSellOrder", "Success").Inc()
			}
		}
	}
}

func generateOrderID() string {
	rand.Seed(uint64(time.Now().UnixNano()))
	return fmt.Sprintf("%d", rand.Intn(1000000)) 
}

func RegisterWebSocket(r *gin.Engine, cfg *config.Config) {
		
	r.GET("/ws/orderbook", func(c *gin.Context) {
		handlers.GetOrderBookWS(c)
	})

	r.GET("/ws/orders", func(c *gin.Context) {
		handlers.WebSocketOrderHandler(c, orderBook)
	})
}

func RegisterRoutes(r *gin.Engine, cfg *config.Config) {
	api := r.Group("/api")

	v1 := api.Group("/v1")
	{
		
		orders := v1.Group("/orders")

		// Health
		health := v1.Group("/health")
		

		// Health
		routers.Health(health)
		
		
		routers.Order(orders, cfg, orderBook)

		r.GET("/metrics", gin.WrapH(promhttp.Handler()))
	}
}

func RegisterValidators() {
	val, ok := binding.Validator.Engine().(*validator.Validate)

	if ok {
		err := val.RegisterValidation("mobile", validation.IranianMobileNumberValidator)
		if err != nil {
			logger.Error(logging.Validation, logging.Startup, err.Error(), nil)
		}

		err = val.RegisterValidation("password", validation.PasswordValidator)
		if err != nil {
			logger.Error(logging.Validation, logging.Startup, err.Error(), nil)
		}
	}
}

func RegisterSwagger(r *gin.Engine, cfg *config.Config) {
	docs.SwaggerInfo.Title = "Order Matcher Microservice"
	docs.SwaggerInfo.Description = "This is the API documentation for the Order Matcher Microservice. The microservice handles matching market and limit orders for various torder-matchingg pairs."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Host = fmt.Sprintf("localhost:%s", cfg.Server.ExternalPort)
	docs.SwaggerInfo.Schemes = []string{"http"}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}


func RegisterPrometheus() {
	err := prometheus.Register(metrics.DbCall)
	if err != nil {
		logger.Error(logging.Prometheus, logging.Startup, err.Error(), nil)
	}

	err = prometheus.Register(metrics.HttpDuration)
	if err != nil {
		logger.Error(logging.Prometheus, logging.Startup, err.Error(), nil)
	}

	err = prometheus.Register(metrics.NewOrder)
	if err != nil {
		logger.Error(logging.Prometheus, logging.Startup, err.Error(), nil)
	}

	err = prometheus.Register(metrics.TradeCompleted)
	if err != nil {
		logger.Error(logging.Prometheus, logging.Startup, err.Error(), nil)
	}

	err = prometheus.Register(metrics.ProcessedOrder)
	if err != nil {
		logger.Error(logging.Prometheus, logging.Startup, err.Error(), nil)
	}

	
	
}
