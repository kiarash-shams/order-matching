# Order Matching Microservice

## Overview
The Order Matching Microservice is designed to handle and process trading orders efficiently in a cryptocurrency exchange platform. It facilitates the matching of buy and sell orders, ensuring that trades are executed in a timely and reliable manner.

## Features
- **Order Matching**: Matches buy and sell orders based on price and time priority.
- **Order Types**: Supports multiple order types including limit orders, market orders, and stop-limit orders.
- **High Performance**: Optimized for high-frequency trading with low latency.
- **Scalability**: Built to handle large volumes of transactions with scalability in mind.
- **WebSocket Integration**: Provides real-time updates of order book changes through WebSocket connections.

## Architecture
The microservice is developed using Go and the Gin framework. It interacts with:
- **Redis**: For storing order book data and facilitating quick access to order information.
- **PostgreSQL**: For persistent storage of completed orders and user transaction history.

## API Endpoints
- `POST /orders` - Place a new order.
- `GET /orders/:id` - Retrieve the details of a specific order.
- `GET /order-book` - Get the current state of the order book.

## WebSocket
To connect to the order book updates via WebSocket, use the following endpoint:
- `ws://your-service-url/ws/order-book?stream=BTC-USD&interval=1000ms`

### Query Parameters
- `stream`: The market stream you want to subscribe to (e.g., BTC-USD).
- `interval`: The interval for sending updates (valid values: `100ms`, `500ms`, `1000ms`).

## Installation
1. Clone the repository:
   ```bash
   git clone https://github.com/your-username/order-matching-microservice.git
   cd order-matching-microservice
