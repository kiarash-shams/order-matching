package main

import (
	"order-matching/api"
	"order-matching/config"
)


func main() {

	// Get Config
	cfg := config.GetConfig()

	// Run Server
	api.InitServer(cfg)
}
