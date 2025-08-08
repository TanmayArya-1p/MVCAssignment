package models

import (
	"errors"
	"inorder/pkg/types"
	"slices"
	"time"
)

var bumpMap map[types.OrderItemStatus]types.OrderItemStatus = map[types.OrderItemStatus]types.OrderItemStatus{
	types.OrderItemStatusPending:   types.OrderItemStatusPreparing,
	types.OrderItemStatusPreparing: types.OrderItemStatusServed,
}

func GetAllOrderedItems() ([]*types.OrderItem, error) {
	rows, err := db.Query("SELECT * FROM order_items ORDER BY issued_at")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []*types.OrderItem = make([]*types.OrderItem, 0)

	for rows.Next() {
		item := &types.OrderItem{}

		var temp []byte
		if err := rows.Scan(&item.ID, &item.OrderID, &item.ItemID, &item.Instructions, &item.Quantity, &item.Price, temp, &item.Status); err != nil {
			return nil, err
		}
		item.IssuedAt, err = time.Parse(time.DateTime, string(temp))
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

func GetOrderedItemByID(id types.OrderItemID) (*types.OrderItem, error) {
	row, err := db.Query("SELECT * FROM order_items WHERE id = ?", id)
	if err != nil {
		return nil, err
	}
	defer row.Close()

	if !row.Next() {
		return nil, errors.New("Order item not found")
	}

	item := &types.OrderItem{}

	var temp []byte
	if err := row.Scan(&item.ID, &item.OrderID, &item.ItemID, &item.Instructions, &item.Quantity, &item.Price, &temp, &item.Status); err != nil {
		return nil, err
	}
	item.IssuedAt, err = time.Parse(time.DateTime, string(temp))
	if err != nil {
		return nil, err
	}
	return item, nil
}

func BumpOrderItemStatus(id types.OrderItemID) (*types.OrderItem, error) {
	item, err := GetOrderedItemByID(id)
	if err != nil {
		return nil, err
	}
	if item.Status == types.OrderItemStatusServed {
		return nil, errors.New("Order item already served")
	}

	newStatus, _ := bumpMap[item.Status]

	_, err = db.Exec("UPDATE order_items SET status = ? WHERE id = ?", newStatus, id)
	if err != nil {
		return nil, err
	}

	item.Status = newStatus
	EvaluateOrderStatus(item.OrderID)
	return item, nil
}

func EvaluateOrderStatus(id types.OrderID) error {
	order, err := GetOrderByID(id)
	if err != nil {
		return err
	}

	if order.Status == types.OrderStatusServed || order.Status == types.OrderStatusPaid {
		return errors.New("Order already served")
	}

	orderedItems, err := GetOrderedItems(order.ID)
	if err != nil {
		return err
	}

	var StatusNames []types.OrderStatus
	for _, item := range orderedItems {
		StatusNames = append(StatusNames, types.OrderStatus(item.Status))
	}
	StatusNames = append(StatusNames, types.OrderStatusServed)
	var minStatus types.OrderStatus = slices.Min(StatusNames)

	if order.Status != minStatus {
		err = UpdateOrder(order, &OrderUpdateInstruction{
			Status: &minStatus,
		})
	}
	return err
}
