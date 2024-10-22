package api

import (
	"fmt"

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
	"github.com/prometheus/client_golang/prometheus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var logger = logging.NewLogger(config.GetConfig())


var orderBook *matchingo.OrderBook

func InitServer(cfg *config.Config) {

	
	gin.SetMode(cfg.Server.RunMode)
	r := gin.New()

	RegisterValidators()
	// RegisterPrometheus()



	orderBook = matchingo.NewOrderBook()
	matchingo.SetDecimalFraction(5) 
	logger.Info(logging.Orderbook, logging.Startup, "Run OrderBook", nil)


	r.Use(middleware.DefaultStructuredLogger(cfg))
	r.Use(middleware.Cors(cfg))
	// r.Use(middleware.Prometheus())

	r.Use(gin.Logger(), gin.CustomRecovery(middleware.ErrorHandler))

	
	RegisterRoutes(r, cfg)

	RegisterWebSocket(r, cfg)
	logger.Info(logging.WebSocket, logging.Startup, "Run WebSockets ", nil)

	RegisterSwagger(r, cfg)

	err := r.Run(fmt.Sprintf(":%s", cfg.Server.InternalPort))
	
	if err != nil {
		logger.Error(logging.General, logging.Startup, err.Error(), nil)
	}
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
	// docs.SwaggerInfo.Host = fmt.Sprintf("matchmaking.liara.run:%s", cfg.Server.ExternalPort)
	docs.SwaggerInfo.Host = "matchmaking.liara.run"
	docs.SwaggerInfo.Schemes = []string{"https"}

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
}
