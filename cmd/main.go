package main

import (
	"inorder/pkg/api"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	log.Println("Starting InOrder Server")
	godotenv.Load()
	api.Serve()

}

//TODO: TEST ORDER, ORDER ITEMS
//TODO: AUTHENTICATION MIDDLEWARE
//TODO: DB MIGRATIONS
//TODO: WRITE CONTROLLERS
//TODO: SERVE CONTROLLERS WITH MUXES
