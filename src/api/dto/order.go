package dto

type OrderRequest struct {
	OrderID     string  `json:"order_id" binding:"required"`    
	OrderType   string  `json:"order_type" binding:"required"`
	Market      string  `json:"market" binding:"required"`
	Price       float64 `json:"price"`
	Amount      float64 `json:"amount" binding:"required"`
	OrderKind   string  `json:"order_kind" binding:"required"`
	
}

type GetOrderRequest struct {
    Market  string `json:"market" binding:"required"`
    OrderID string `json:"order_id" binding:"required"`
}


type OrderResponse struct {
    OrderID  string  `json:"orderID"`
    Role     string  `json:"role"`
    IsQuote  bool    `json:"isQuote"`
    Price    string  `json:"price"`
    Quantity string  `json:"quantity"`
}

type CancelOrderRequest struct {
    Market  string `json:"market" binding:"required"`
    OrderID string `json:"order_id" binding:"required"`
}

type GetOrderBookRequest struct {
	Market string `json:"market" binding:"required"`
}