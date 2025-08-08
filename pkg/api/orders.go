package api

import (
	"inorder/pkg/controllers"
	"inorder/pkg/middleware"
	"inorder/pkg/types"
	"net/http"

	"github.com/gorilla/mux"
)

func SetupOrdersRoutes(root *mux.Router) {
	var orderRouter = root.PathPrefix("/api/orders").Subrouter()
	orderRouter.Use(middleware.AuthenticationMiddleware(false))
	orderRouter.Handle("", middleware.AuthorizationMiddleware(types.ChefRole)(http.HandlerFunc(controllers.GetAllOrdersController))).Methods("GET")
	orderRouter.Handle("", http.HandlerFunc(controllers.CreateOrderController)).Methods("POST")
	orderRouter.Handle("/{orderid}", middleware.AuthorizationMiddleware(types.ChefRole)(http.HandlerFunc(controllers.DeleteOrderController))).Methods("DELETE")
	orderRouter.Handle("/{orderid}", http.HandlerFunc(controllers.GetOrderController)).Methods("GET")
	orderRouter.Handle("/{orderid}", middleware.AuthorizationMiddleware(types.ChefRole)(http.HandlerFunc(controllers.UpdateOrderController))).Methods("PUT")
	orderRouter.Handle("/{orderid}/items", http.HandlerFunc(controllers.GetAllOrderItemsController)).Methods("GET")
	orderRouter.Handle("/{orderid}/items", http.HandlerFunc(controllers.OrderNewItemController)).Methods("POST")
	orderRouter.Handle("/{orderid}/bill", http.HandlerFunc(controllers.GetOrderBillController)).Methods("GET")
	orderRouter.Handle("/{orderid}/bill/pay", middleware.AuthorizationMiddleware(types.ChefRole)(http.HandlerFunc(controllers.PayOrderController))).Methods("PUT")
	orderRouter.Handle("/item/{itemid}/bump", middleware.AuthorizationMiddleware(types.ChefRole)(http.HandlerFunc(controllers.BumpOrderItemStatusController))).Methods("POST")
	orderRouter.Handle("/my", http.HandlerFunc(controllers.GetUserOrdersController)).Methods("GET")
}
