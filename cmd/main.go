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
