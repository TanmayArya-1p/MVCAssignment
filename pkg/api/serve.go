package api

import (
	"inorder/pkg/config"
	"inorder/pkg/controllers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Serve() error {
	var root *mux.Router = mux.NewRouter()

	root.HandleFunc("/openapi.json", controllers.OpenAPISpec).Methods("GET")

	SetupAuthRoutes(root)
	SetupUserRoutes(root)
	SetupOrdersRoutes(root)
	SetupItemsRoutes(root)
	SetupStaticFileRoutes(root)

	log.Println("Serving HTTP Server on Port", config.Config.InOrder.PORT)
	err := http.ListenAndServe(":"+config.Config.InOrder.PORT, root)
	if err != nil {
		log.Fatal(err)
	}
	return err
}
