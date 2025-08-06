package api

import (
	"inorder/pkg/config"
	"inorder/pkg/controllers"
	"inorder/pkg/middleware"
	"inorder/pkg/types"
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
	authRouter.Handle("/verify", middleware.AuthenticationMiddleware(true)(http.HandlerFunc(controllers.VerifyController))).Methods("GET")
	authRouter.Handle("/refresh/lmao", middleware.AuthenticationMiddleware(true)(http.HandlerFunc(controllers.RefreshController))).Methods("GET")
	authRouter.Handle("/logout", middleware.AuthenticationMiddleware(true)(http.HandlerFunc(controllers.LogoutController))).Methods("POST")

	var userRouter *mux.Router = root.PathPrefix("/api/users").Subrouter()
	userRouter.Use(middleware.AuthenticationMiddleware(false))
	userRouter.Use(middleware.AuthorizationMiddleware(types.AdminRole))
	userRouter.HandleFunc("", controllers.GetAllUsersController).Methods("GET")
	userRouter.HandleFunc("", controllers.CreateUserController).Methods("POST")
	userRouter.HandleFunc("/{userid}", controllers.DeleteUserController).Methods("DELETE")
	userRouter.HandleFunc("/{userid}", controllers.GetUserByIDController).Methods("GET")
	userRouter.HandleFunc("/{userid}", controllers.UpdateUserController).Methods("PUT")

	var itemRouter *mux.Router = root.PathPrefix("/api/items").Subrouter()
	itemRouter.Use(middleware.AuthenticationMiddleware(false))
	itemRouter.HandleFunc("", controllers.GetAllItemsController).Methods("GET")
	itemRouter.Handle("", middleware.AuthorizationMiddleware(types.ChefRole)(http.HandlerFunc(controllers.CreateItemController))).Methods("POST")
	itemRouter.HandleFunc("/tags", controllers.GetAllTagsController).Methods("GET")
	itemRouter.Handle("/{itemid}", middleware.AuthorizationMiddleware(types.ChefRole)(http.HandlerFunc(controllers.DeleteItemController))).Methods("DELETE")
	itemRouter.HandleFunc("/{itemid}", controllers.GetItemByIDController).Methods("GET")
	itemRouter.Handle("/{itemid}", middleware.AuthorizationMiddleware(types.ChefRole)(http.HandlerFunc(controllers.UpdateItemController))).Methods("PUT")
	itemRouter.Handle("/upload", middleware.AuthorizationMiddleware(types.ChefRole)(http.HandlerFunc(controllers.UploadImageController))).Methods("POST")

	root.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))

	log.Println("Serving HTTP Server on Port", config.Config.InOrder.PORT)
	err := http.ListenAndServe(":"+config.Config.InOrder.PORT, root)
	if err != nil {
		log.Fatal(err)
	}
	return err
}
