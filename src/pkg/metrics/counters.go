package metrics

import "github.com/prometheus/client_golang/prometheus"



var DbCall = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "db_calls_total",
		Help: "Number of database calls",
	},[]string{"type_name","operation_name", "status"},
)

// Counter for new orders
var NewOrder = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "new_orders_total",
		Help: "Number of new orders created",
	}, 
	[]string{"order_type", "status"},
)


// Counter for completed trades (orders that have been executed)
var TradeCompleted = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "completed_trades_total",
		Help: "Number of completed trades",
	}, 
	[]string{"trade_type", "status"},
)


// Counter for processed orders in orderBook
var ProcessedOrder = prometheus.NewCounterVec(
    prometheus.CounterOpts{
        Name: "processed_orders_total",
        Help: "Number of orders processed by orderBook",
    },
    []string{"order_type", "status"}, // Labels for categorizing metrics
)