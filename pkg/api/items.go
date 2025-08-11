package api

import (
	"inorder/pkg/controllers"
	"inorder/pkg/middleware"
	"inorder/pkg/types"
	"net/http"

	"github.com/gorilla/mux"
)

func SetupItemsRoutes(root *mux.Router) {
	var itemRouter *mux.Router = root.PathPrefix("/api/items").Subrouter()
	chefAuth := middleware.AuthorizationMiddleware(types.ChefRole)

	itemRouter.Use(middleware.AuthenticationMiddleware(false))
	itemRouter.HandleFunc("", controllers.GetAllItemsController).Methods("GET")                         // GET /api/items -> Get all items in the menu
	itemRouter.Handle("", chefAuth(http.HandlerFunc(controllers.CreateItemController))).Methods("POST") // POST /api/items -> Create and add a new item to the menu
	itemRouter.HandleFunc("/tags", controllers.GetAllTagsController).Methods("GET")
	itemRouter.HandleFunc("/bytags", controllers.GetItemsOfTagsController).Methods("GET")
	itemRouter.Handle("/{itemid}", chefAuth(http.HandlerFunc(controllers.DeleteItemController))).Methods("DELETE")
	itemRouter.HandleFunc("/{itemid}", controllers.GetItemByIDController).Methods("GET")
	itemRouter.Handle("/{itemid}", chefAuth(http.HandlerFunc(controllers.UpdateItemController))).Methods("PUT")
	itemRouter.Handle("/upload", chefAuth(http.HandlerFunc(controllers.UploadImageController))).Methods("POST")
}
