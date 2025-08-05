package models

import (
	"errors"
	"inorder/pkg/types"
	"time"
)

func GetOrderedItems(orderID types.OrderID) ([]*types.OrderItem, error) {
	rows, err := db.Query("SELECT * FROM order_items WHERE order_id = ?", orderID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var otpt []*types.OrderItem
	if exists := rows.Next(); !exists {
		return otpt, nil
	}

	for {
		var item types.OrderItem
		var temp []uint8

		if err := rows.Scan(&item.ID, &item.OrderID, &item.ItemID, &item.Instructions, &item.Quantity, &item.Price, &temp, &item.Status); err != nil {
			return nil, err
		}
		item.IssuedAt, err = time.Parse(time.DateTime, string(temp))

		if err != nil {
			return nil, err
		}

		otpt = append(otpt, &item)

		if next := rows.Next(); !next {
			break
		}

	}

	return otpt, nil
}

func GetOrderByID(orderID types.OrderID) (*types.Order, error) {
	var order types.Order

	var issuedAtTemp []uint8
	var paidAtTemp []uint8

	err := db.QueryRow("SELECT * FROM orders WHERE id = ?", orderID).Scan(&order.ID, &order.IssuedBy, &issuedAtTemp, &order.Status, &order.BillableAmount, &order.TableNo, &order.Waiter, &paidAtTemp, &order.Tip)
	if err != nil {
		return nil, err
	}
	order.IssuedAt, err = time.Parse(time.DateTime, string(issuedAtTemp))
	if err != nil {
		return nil, err
	}
	order.PaidAt, err = time.Parse(time.DateTime, string(paidAtTemp))
	if err != nil {
		return nil, err
	}

	return &order, nil
}

func CreateOrder(order *types.Order) (types.OrderID, error) {
	res, err := db.Exec("INSERT INTO orders (issued_by, table_no) VALUES (?, ?)", order.IssuedBy, order.TableNo)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	order.ID = types.OrderID(id)
	return order.ID, nil
}

func DeleteAllItemOrders(orderID types.OrderID) error {
	_, err := db.Exec("DELETE FROM order_items WHERE order_id = ?", orderID)
	return err
}

func DeleteOrder(orderID types.OrderID) error {
	err := DeleteAllItemOrders(orderID)
	if err != nil {
		return err
	}
	_, err = db.Exec("DELETE FROM orders WHERE id = ?", orderID)
	return err
}

func GetAllOrders(page types.Page) ([]*types.Order, error) {
	rows, err := db.Query("SELECT * FROM orders LIMIT ? OFFSET ?", page.Limit, page.Offset)

	var otpt []*types.Order
	if err != nil {
		return otpt, err
	}
	defer rows.Close()

	if exists := rows.Next(); !exists {
		return otpt, err
	}
	for {
		var curr types.Order
		var issuedAtTemp, paidAtTemp []uint8
		err := rows.Scan(&curr.ID, &curr.IssuedBy, &issuedAtTemp, &curr.Status, &curr.BillableAmount, &curr.TableNo, &curr.Waiter, &paidAtTemp, &curr.Tip)
		if err != nil {
			return otpt, err
		}
		curr.IssuedAt, err = time.Parse(time.DateTime, string(issuedAtTemp))
		if err != nil {
			return otpt, err
		}
		curr.PaidAt, err = time.Parse(time.DateTime, string(paidAtTemp))
		if err != nil {
			return otpt, err
		}
		otpt = append(otpt, &curr)

		if isNext := rows.Next(); !isNext {
			break
		}
	}
	return otpt, nil
}

func PayBill(order types.Order, waiter types.UserID, tip float32) error {
	if order.Status != types.OrderStatusBilled {
		return errors.New("Order is not billed yet")
	}
	_, err := db.Exec("UPDATE orders SET status= ?,waiter = ?, tip = ?, paid_at = NOW() WHERE id = ?", types.OrderStatusPaid, waiter, tip, order.ID)
	if err != nil {
		return err
	}
	order.Waiter = waiter
	order.PaidAt = time.Now()
	order.Tip = tip
	return nil
}

type OrderUpdateInstruction struct {
	Status         *types.OrderStatus
	TableNo        *types.TableID
	BillableAmount *float32
}

func UpdateOrder(order *types.Order, instruction *OrderUpdateInstruction) error {

	if instruction.BillableAmount == nil {
		instruction.BillableAmount = &order.BillableAmount
	}
	if instruction.Status == nil {
		instruction.Status = &order.Status
	}
	if instruction.TableNo == nil {
		instruction.TableNo = &order.TableNo
	}
	_, err := db.Exec("UPDATE orders SET status = ?, table_no = ?, billable_amount = ? WHERE id = ?", instruction.Status, instruction.TableNo, instruction.BillableAmount, order.ID)
	if err != nil {
		return err
	}
	order.Status = *instruction.Status
	order.TableNo = *instruction.TableNo
	order.BillableAmount = *instruction.BillableAmount
	return nil
}

func ResolveBillableAmount(order *types.Order, toBill bool) error {
	items, err := GetOrderedItems(order.ID)
	if err != nil {
		return err
	}

	var amt float32
	for _, item := range items {
		amt += item.Price
	}

	var changeStatus *types.OrderStatus = nil
	if toBill {
		changeStatus = &types.OrderStatusBilled
	}

	return UpdateOrder(order, &OrderUpdateInstruction{
		Status:         changeStatus,
		BillableAmount: &amt,
	})
}

func OrderNewItem(order *types.Order, itemID types.ItemID, quantity int, instructions string) error {
	if quantity < 1 {
		return errors.New("Quantity must be greater than or equal to 1")
	}
	if order.Status == types.OrderStatusBilled || order.Status == types.OrderStatusPaid {
		return errors.New("Order status not open to item orders")
	}

	item, err := GetItemByID(itemID)
	if err != nil {
		return err
	}

	var price float32 = float32(item.Price) * float32(quantity)

	_, err = db.Exec("INSERT INTO order_items (order_id, item_id, quantity, price,instructions) VALUES (?, ?, ?, ?, ?)", order.ID, item.ID, quantity, price, instructions)
	if err != nil {
		return err
	}
	UpdateOrder(order, &OrderUpdateInstruction{
		Status: &types.OrderStatusPending,
	})
	return nil
}
