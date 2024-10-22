package logging

type Category string
type SubCategory string
type ExtraKey string

const (
	General         Category = "General"
	IO              Category = "IO"
	Internal        Category = "Internal"
	Postgres        Category = "Postgres"
	Redis           Category = "Redis"
	Validation      Category = "Validation"
	RequestResponse Category = "RequestResponse"
	Prometheus      Category = "Prometheus"
	Orderbook       Category = "Orderbook"
	WebSocket       Category = "WebSocket"
	
)

const (
	// General
	Startup         SubCategory = "Startup"
	ExternalService SubCategory = "ExternalService"

	// Postgres
	Migration SubCategory = "Migration"
	Select    SubCategory = "Select"
	Rollback  SubCategory = "Rollback"
	Update    SubCategory = "Update"
	Delete    SubCategory = "Delete"
	Insert    SubCategory = "Insert"

	// Internal
	Api                 SubCategory = "Api"
	HashPassword        SubCategory = "HashPassword"
	DefaultRoleNotFound SubCategory = "DefaultRoleNotFound"
	FailedToCreateUser  SubCategory = "FailedToCreateUser"

	// Validation
	MobileValidation   SubCategory = "MobileValidation"
	PasswordValidation SubCategory = "PasswordValidation"

	// IO
	RemoveFile SubCategory = "RemoveFile"

	// Orderbook
	Process          	 SubCategory = "Process"          // Log when an order is placed
	OrderPlaced          SubCategory = "OrderPlaced"          // Log when an order is placed
	OrderExecuted        SubCategory = "OrderExecuted"        // Log when an order is executed
	OrderCancelled       SubCategory = "OrderCancelled"       // Log when an order is cancelled
	OrderUpdate          SubCategory = "OrderUpdate"          // Log when an order is updated
	InsufficientFunds    SubCategory = "InsufficientFunds"    // Log when there are insufficient funds to place an order
	OrderNotFound        SubCategory = "OrderNotFound"        // Log when an order is not found
	MarketNotFound       SubCategory = "MarketNotFound"        // Log when an order is not found
	OrderPriceAdjustment SubCategory = "OrderPriceAdjustment" // Log when an order price is adjusted
	NewLimitOrder 		 SubCategory = "NewLimitOrder" // Log when an order price is adjusted
	NewMarketOrder 	 	 SubCategory = "NewMarketOrder" // Log when an order price is adjusted
	GetOrderId 	 	 SubCategory = "GetOrderId" // Log when an order price is adjusted
	ListOrderBook 	 	 SubCategory = "ListOrderBook" // Log when an order price is adjusted

	// WebSocket
	ConnectionOpened   SubCategory = "ConnectionOpened"   // Log when a WebSocket connection is opened
	ConnectionClosed   SubCategory = "ConnectionClosed"   // Log when a WebSocket connection is closed
	MessageReceived    SubCategory = "MessageReceived"    // Log when a message is received
	MessageSent        SubCategory = "MessageSent"        // Log when a message is sent
	ErrorOccurred      SubCategory = "ErrorOccurred"      // Log when an error occurs
	InvalidMessage     SubCategory = "InvalidMessage"     // Log when an invalid message is received
	ReconnectAttempted SubCategory = "ReconnectAttempted" // Log when a reconnection is attempted

	
)


const (
	AppName      ExtraKey = "AppName"
	LoggerName   ExtraKey = "Logger"
	ClientIp     ExtraKey = "ClientIp"
	HostIp       ExtraKey = "HostIp"
	Method       ExtraKey = "Method"
	StatusCode   ExtraKey = "StatusCode"
	BodySize     ExtraKey = "BodySize"
	Path         ExtraKey = "Path"
	Latency      ExtraKey = "Latency"
	RequestBody  ExtraKey = "RequestBody"
	ResponseBody ExtraKey = "ResponseBody"
	ErrorMessage ExtraKey = "ErrorMessage"

	// WebSocket specific extra keys
	MessageID ExtraKey = "MessageID" // New key for the ID of the message
	Channel   ExtraKey = "Channel"    // New key for the channel/topic of the WebSocket
)