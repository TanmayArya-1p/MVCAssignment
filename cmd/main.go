package main

import (
	"inorder/pkg/api"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	log.Println("Starting InOrder Server")
	godotenv.Load()
	// _, err := workers.StartCleanupWorker()
	// if err != nil {
	// 	log.Fatalf("Failed to start cleanup worker: %v", err)
	// }
	api.Serve()
}
