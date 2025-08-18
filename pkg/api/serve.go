package api

import (
	"inorder/pkg/config"
	"inorder/pkg/controllers"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func Serve() error {
	var root *mux.Router = mux.NewRouter()

	root.HandleFunc("/openapi.json", controllers.OpenAPISpec).Methods("GET")

	SetupAuthRoutes(root)
	SetupUserRoutes(root)
	SetupOrdersRoutes(root)
	SetupItemsRoutes(root)
	SetupStaticFileRoutes(root)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})

	var port string
	if os.Getenv("INORDER_PORT") == "" {
		port = config.Config.InOrder.PORT
	} else {
		port = os.Getenv("INORDER_PORT")
	}

	log.Println("Serving HTTP Server on Port", port)

	err := http.ListenAndServe(":"+port, c.Handler(root))
	if err != nil {
		log.Fatal(err)
	}
	return err
}
