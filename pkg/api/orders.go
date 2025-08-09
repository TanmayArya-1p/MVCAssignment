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
	chefAuth := middleware.AuthorizationMiddleware(types.ChefRole)

	orderRouter.Use(middleware.AuthenticationMiddleware(false))
	orderRouter.Handle("", chefAuth(http.HandlerFunc(controllers.GetAllOrdersController))).Methods("GET") // GET /api/orders -> Get all orders
	orderRouter.Handle("", http.HandlerFunc(controllers.CreateOrderController)).Methods("POST")           // POST /api/orders -> Create a new order
	orderRouter.Handle("/my", http.HandlerFunc(controllers.GetUserOrdersController)).Methods("GET")
	orderRouter.Handle("/item/{itemid}/bump", chefAuth(http.HandlerFunc(controllers.BumpOrderItemStatusController))).Methods("POST")
	orderRouter.Handle("/{orderid}", chefAuth(http.HandlerFunc(controllers.DeleteOrderController))).Methods("DELETE")
	orderRouter.Handle("/{orderid}", http.HandlerFunc(controllers.GetOrderController)).Methods("GET")
	orderRouter.Handle("/{orderid}", chefAuth(http.HandlerFunc(controllers.UpdateOrderController))).Methods("PUT")
	orderRouter.Handle("/{orderid}/items", http.HandlerFunc(controllers.GetAllOrderItemsController)).Methods("GET")
	orderRouter.Handle("/{orderid}/items", http.HandlerFunc(controllers.OrderNewItemController)).Methods("POST")
	orderRouter.Handle("/{orderid}/bill", http.HandlerFunc(controllers.GetOrderBillController)).Methods("GET")
	orderRouter.Handle("/{orderid}/bill/pay", chefAuth(http.HandlerFunc(controllers.PayOrderController))).Methods("POST")
}
