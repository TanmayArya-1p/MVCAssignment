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
	pg, err := utils.Paginate(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
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
	user := r.Context().Value(types.UserContextKey).(*types.User)

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

	if user.Role == types.UserRole && order.IssuedBy != user.ID {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	json.NewEncoder(w).Encode(order)
}

func GetUserOrdersController(w http.ResponseWriter, r *http.Request) {
	pg, err := utils.Paginate(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user := r.Context().Value(types.UserContextKey).(*types.User)
	userOrders, err := models.GetAllOrdersByUser(user, &pg)

	if err != nil {
		http.Error(w, "Internal server error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(userOrders)
}

func CreateOrderController(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(types.UserContextKey).(*types.User)

	var body CreateOrderRequest
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	tableNo := body.TableNo
	if !(tableNo <= 100 && tableNo > 0) {
		http.Error(w, "Invalid table_no parameter", http.StatusBadRequest)
		return
	}

	var order types.Order
	order.IssuedBy = user.ID
	order.TableNo = types.TableID(tableNo)

	_, err := models.CreateOrder(&order)
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
	user := r.Context().Value(types.UserContextKey).(*types.User)

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

	if user.Role == types.UserRole && user.ID != order.IssuedBy {
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
	user := r.Context().Value(types.UserContextKey).(*types.User)

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

	if user.Role == types.UserRole && user.ID != order.IssuedBy {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	orderItemID, err := models.OrderNewItem(order, body.ItemID, body.Quantity, body.Instructions)
	if err != nil {
		http.Error(w, "Internal server error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]any{"message": "Item ordered successfully", "order_item_id": orderItemID})
}

func GetOrderBillController(w http.ResponseWriter, r *http.Request) {
	tempID, ok := mux.Vars(r)["orderid"]
	user := r.Context().Value(types.UserContextKey).(*types.User)

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

	if user.Role == types.UserRole && user.ID != order.IssuedBy {
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
	user := r.Context().Value(types.UserContextKey).(*types.User)

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
	amount64, ok := body["amount"].(float64)
	amount := float32(amount64)
	if !ok {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	if amount <= 0 || amount < order.BillableAmount {
		http.Error(w, "Invalid amount", http.StatusBadRequest)
		return
	}

	err = models.PayBill(order, user.ID, amount-order.BillableAmount)
	if err != nil {
		http.Error(w, "Internal server error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]any{"message": "Order paid successfully", "order": order})
}
