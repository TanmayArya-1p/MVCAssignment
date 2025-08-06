package api

import (
	"inorder/pkg/config"
	"inorder/pkg/controllers"
	"inorder/pkg/middleware"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Serve() error {
	var root *mux.Router = mux.NewRouter()

	root.HandleFunc("/openapi.json", controllers.OpenAPISpec).Methods("GET")

	var authRouter *mux.Router = root.PathPrefix("/api/auth").Subrouter()
	authRouter.HandleFunc("/register", controllers.RegisterController).Methods("POST")
	authRouter.HandleFunc("/login", controllers.LoginController).Methods("POST")
	authRouter.Handle("/verify", middleware.AuthenticationMiddleware(http.HandlerFunc(controllers.VerifyController), true)).Methods("GET")
	authRouter.Handle("/refresh", middleware.AuthenticationMiddleware(http.HandlerFunc(controllers.RefreshController), true)).Methods("GET")
	authRouter.Handle("/logout", middleware.AuthenticationMiddleware(http.HandlerFunc(controllers.LogoutController), true)).Methods("POST")

	log.Println("Serving HTTP Server on Port", config.Config.InOrder.PORT)
	err := http.ListenAndServe(":"+config.Config.InOrder.PORT, root)
	if err != nil {
		log.Fatal(err)
	}
	return err
}
