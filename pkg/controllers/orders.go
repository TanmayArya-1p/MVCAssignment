package controllers

import (
	"encoding/json"
	"inorder/pkg/models"
	"inorder/pkg/types"
	"inorder/pkg/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func BumpOrderItemStatusController(w http.ResponseWriter, r *http.Request) {
	tempID, ok := mux.Vars(r)["itemid"]

	orderItemID, err := strconv.Atoi(tempID)
	if err != nil {
		http.Error(w, "Invalid order item id", http.StatusBadRequest)
		return
	}

	if !ok {
		http.Error(w, "Invalid order item id", http.StatusBadRequest)
		return
	}

	item, err := models.BumpOrderItemStatus(types.OrderItemID(orderItemID))
	if err != nil {
		http.Error(w, "Failed to bump order item status: "+err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Bumped status to " + string(item.Status)})
}

func GetAllOrdersController(w http.ResponseWriter, r *http.Request) {
	limit := r.URL.Query().Get("limit")
	offset := r.URL.Query().Get("offset")

	var pg types.Page
	var err error

	if limit == "" {
		pg.Limit = 10
	} else {
		pg.Limit, err = strconv.Atoi(limit)
		if err != nil || pg.Limit <= 0 {
			http.Error(w, "Invalid limit parameter", http.StatusBadRequest)
			return
		}
	}

	if offset == "" {
		pg.Offset = types.DefaultOffset
	} else {
		pg.Offset, err = strconv.Atoi(offset)
		if err != nil || pg.Offset < 0 {
			http.Error(w, "Invalid offset parameter", http.StatusBadRequest)
			return
		}
	}

	orders, err := models.GetAllOrders(pg)
	if err != nil {
		http.Error(w, "Internal server error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(orders)
}

func GetOrderController(w http.ResponseWriter, r *http.Request) {
	tempID, ok := mux.Vars(r)["orderid"]
	if !ok {
		http.Error(w, "Invalid id parameter", http.StatusBadRequest)
		return
	}
	user := r.Context().Value("user").(*types.User)

	orderIDx, err := strconv.Atoi(tempID)
	if err != nil {
		http.Error(w, "Invalid id parameter", http.StatusBadRequest)
		return
	}

	orderID := types.OrderID(orderIDx)
	order, err := models.GetOrderByID(orderID)
	if err != nil {
		if err == utils.ErrOrderNotFound {
			http.Error(w, "Order not found", http.StatusNotFound)
		} else {
			http.Error(w, "Internal server error: "+err.Error(), http.StatusInternalServerError)
		}
		return
	}

	if order.IssuedBy != user.ID {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	json.NewEncoder(w).Encode(order)
}

func GetUserOrdersController(w http.ResponseWriter, r *http.Request) {
	limit := r.URL.Query().Get("limit")
	offset := r.URL.Query().Get("offset")

	var pg types.Page
	var err error

	if limit == "" {
		pg.Limit = 10
	} else {
		pg.Limit, err = strconv.Atoi(limit)
		if err != nil || pg.Limit <= 0 {
			http.Error(w, "Invalid limit parameter", http.StatusBadRequest)
			return
		}
	}

	if offset == "" {
		pg.Offset = types.DefaultOffset
	} else {
		pg.Offset, err = strconv.Atoi(offset)
		if err != nil || pg.Offset < 0 {
			http.Error(w, "Invalid offset parameter", http.StatusBadRequest)
			return
		}
	}
	//TODO: PAGIFY UTILS

	user := r.Context().Value("user").(*types.User)

	userOrders, err := models.GetAllOrdersByUser(user, &pg)
	if err != nil {
		http.Error(w, "Internal server error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(userOrders)
}

func CreateOrderController(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(*types.User)

	var body map[string]string
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if body["table_no"] == "" {
		http.Error(w, "Missing table_no parameter", http.StatusBadRequest)
		return
	}
	tableNo, err := strconv.Atoi(body["table_no"])
	if err != nil || !(tableNo <= 100 && tableNo >= 0) {
		http.Error(w, "Invalid table_no parameter", http.StatusBadRequest)
		return
	}

	var order types.Order
	order.IssuedBy = user.ID
	order.TableNo = types.TableID(tableNo)

	//TODO: VERIFY CREATE MODELS FOR ALL TABLES WITH DEFAULT VALUES IN SCHEMA

	_, err = models.CreateOrder(&order)
	if err != nil {
		http.Error(w, "Internal server error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(order)
}

func DeleteOrderController(w http.ResponseWriter, r *http.Request) {
	tempID, ok := mux.Vars(r)["orderid"]
	if !ok {
		http.Error(w, "Invalid id parameter", http.StatusBadRequest)
		return
	}
	orderIDx, err := strconv.Atoi(tempID)
	if err != nil {
		http.Error(w, "Invalid id parameter", http.StatusBadRequest)
		return
	}

	orderID := types.OrderID(orderIDx)
	order, err := models.GetOrderByID(orderID)
	if err != nil {
		if err == utils.ErrOrderNotFound {
			http.Error(w, "Order not found", http.StatusNotFound)
		} else {
			http.Error(w, "Internal server error: "+err.Error(), http.StatusInternalServerError)
		}
		return
	}

	err = models.DeleteOrder(orderID)
	if err != nil {
		http.Error(w, "Internal server error: "+err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]any{"message": "Order deleted successfully", "order": order})
}

func UpdateOrderController(w http.ResponseWriter, r *http.Request) {
	tempID, ok := mux.Vars(r)["orderid"]
	if !ok {
		http.Error(w, "Invalid id parameter", http.StatusBadRequest)
		return
	}
	orderID, err := strconv.Atoi(tempID)
	if err != nil {
		http.Error(w, "Invalid id parameter", http.StatusBadRequest)
		return
	}

	var body models.OrderUpdateInstruction
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	//TODO: INTERNAL SERVER ERROR UTIL FUNCTIONS

	order, err := models.GetOrderByID(types.OrderID(orderID))
	if err != nil {
		if err == utils.ErrOrderNotFound {
			http.Error(w, "Order not found", http.StatusNotFound)
		} else {
			http.Error(w, "Internal server error: "+err.Error(), http.StatusInternalServerError)
		}
		return
	}
	err = models.UpdateOrder(order, &body)
	if err != nil {
		http.Error(w, "Internal server error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	order, err = models.GetOrderByID(types.OrderID(orderID))
	if err != nil {
		if err == utils.ErrOrderNotFound {
			http.Error(w, "Order not found", http.StatusNotFound)
		} else {
			http.Error(w, "Internal server error: "+err.Error(), http.StatusInternalServerError)
		}
		return
	}

	json.NewEncoder(w).Encode(map[string]any{"message": "Order updated successfully", "order": order})
}

func GetAllOrderItemsController(w http.ResponseWriter, r *http.Request) {
	tempID, ok := mux.Vars(r)["orderid"]
	user := r.Context().Value("user").(*types.User)

	if !ok {
		http.Error(w, "Invalid id parameter", http.StatusBadRequest)
		return
	}
	orderID, err := strconv.Atoi(tempID)
	if err != nil {
		http.Error(w, "Invalid id parameter", http.StatusBadRequest)
		return
	}
	order, err := models.GetOrderByID(types.OrderID(orderID))
	if err != nil {
		if err == utils.ErrOrderNotFound {
			http.Error(w, "Order not found", http.StatusNotFound)
		} else {
			http.Error(w, "Internal server error: "+err.Error(), http.StatusInternalServerError)
		}
		return
	}

	if order.IssuedBy != user.ID {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	orderItems, err := models.GetOrderedItems(types.OrderID(orderID))
	if err != nil {
		http.Error(w, "Internal server error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(orderItems)
}

func OrderNewItemController(w http.ResponseWriter, r *http.Request) {
	tempID, ok := mux.Vars(r)["orderid"]
	user := r.Context().Value("user").(*types.User)

	if !ok {
		http.Error(w, "Invalid id parameter", http.StatusBadRequest)
		return
	}
	orderID, err := strconv.Atoi(tempID)
	if err != nil {
		http.Error(w, "Invalid id parameter", http.StatusBadRequest)
		return
	}

	var body OrderItemCRUDRequest
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	order, err := models.GetOrderByID(types.OrderID(orderID))
	if err != nil {
		if err == utils.ErrOrderNotFound {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, "Internal server error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if order.IssuedBy != user.ID {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	err = models.OrderNewItem(order, body.ItemID, body.Quantity, body.Instructions)
	if err != nil {
		http.Error(w, "Internal server error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]any{"message": "Item ordered successfully"})
}

func GetOrderBillController(w http.ResponseWriter, r *http.Request) {
	tempID, ok := mux.Vars(r)["orderid"]
	user := r.Context().Value("user").(*types.User)

	//TODO: MAKE CONTEXT KEY A CONST
	if !ok {
		http.Error(w, "Invalid id parameter", http.StatusBadRequest)
		return
	}
	orderID, err := strconv.Atoi(tempID)
	if err != nil {
		http.Error(w, "Invalid id parameter", http.StatusBadRequest)
		return
	}

	order, err := models.GetOrderByID(types.OrderID(orderID))
	if err != nil {
		if err == utils.ErrOrderNotFound {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, "Internal server error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if user.ID != order.IssuedBy {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	resolveStat := r.URL.Query().Get("resolve")
	if resolveStat == "true" {
		err = models.ResolveBillableAmount(order, true)
	} else {
		err = models.ResolveBillableAmount(order, false)
	}
	if err != nil {
		http.Error(w, "Internal server error: "+err.Error(), http.StatusInternalServerError)
		return
	}
	orderItems, err := models.GetOrderedItems(order.ID)
	if err != nil {
		http.Error(w, "Internal server error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]any{"items": orderItems, "status": order.Status, "order_id": order.ID, "billable_amount": order.BillableAmount})
}

func PayOrderController(w http.ResponseWriter, r *http.Request) {
	tempID, ok := mux.Vars(r)["orderid"]
	user := r.Context().Value("user").(*types.User)

	if !ok {
		http.Error(w, "Invalid id parameter", http.StatusBadRequest)
		return
	}
	orderID, err := strconv.Atoi(tempID)
	if err != nil {
		http.Error(w, "Invalid id parameter", http.StatusBadRequest)
		return
	}

	order, err := models.GetOrderByID(types.OrderID(orderID))
	if err != nil {
		if err == utils.ErrOrderNotFound {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, "Internal server error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	var body map[string]any
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	amount := body["amount"].(float32)

	err = models.PayBill(order, user.ID, amount)
	if err != nil {
		http.Error(w, "Internal server error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]any{"message": "Order paid successfully", "order": order})
}
