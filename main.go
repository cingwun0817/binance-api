package main

import (
	"binance-api/internal/common"
	"binance-api/routes"
	"log"
)

func main() {
	// init
	common.LoadConfig()
	common.Connect()

	// route
	app := routes.New()

	// start the server on port 3001
	log.Fatal(app.Listen(":3001"))
}
