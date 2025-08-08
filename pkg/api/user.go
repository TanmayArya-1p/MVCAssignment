package api

import (
	"inorder/pkg/controllers"
	"inorder/pkg/middleware"
	"inorder/pkg/types"

	"github.com/gorilla/mux"
)

func SetupUserRoutes(root *mux.Router) {
	var userRouter *mux.Router = root.PathPrefix("/api/users").Subrouter()
	userRouter.Use(middleware.AuthenticationMiddleware(false))
	userRouter.Use(middleware.AuthorizationMiddleware(types.AdminRole))
	userRouter.HandleFunc("", controllers.GetAllUsersController).Methods("GET")
	userRouter.HandleFunc("", controllers.CreateUserController).Methods("POST")
	userRouter.HandleFunc("/{userid}", controllers.DeleteUserController).Methods("DELETE")
	userRouter.HandleFunc("/{userid}", controllers.GetUserByIDController).Methods("GET")
	userRouter.HandleFunc("/{userid}", controllers.UpdateUserController).Methods("PUT")
}
