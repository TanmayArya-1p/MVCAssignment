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

//TODO: MAYBE TRY CACHING
//TODO: STRESS TESTING
//TODO: DOCKERISE
