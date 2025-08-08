package api

import (
	"inorder/pkg/controllers"
	"inorder/pkg/middleware"
	"net/http"

	"github.com/gorilla/mux"
)

func SetupAuthRoutes(root *mux.Router) {
	var authRouter *mux.Router = root.PathPrefix("/api/auth").Subrouter()
	authRouter.HandleFunc("/register", controllers.RegisterController).Methods("POST")
	authRouter.HandleFunc("/login", controllers.LoginController).Methods("POST")
	authRouter.Handle("/verify", middleware.AuthenticationMiddleware(true)(http.HandlerFunc(controllers.VerifyController))).Methods("GET")
	authRouter.Handle("/refresh/lmao", middleware.AuthenticationMiddleware(true)(http.HandlerFunc(controllers.RefreshController))).Methods("GET")
	authRouter.Handle("/logout", middleware.AuthenticationMiddleware(true)(http.HandlerFunc(controllers.LogoutController))).Methods("POST")
}
