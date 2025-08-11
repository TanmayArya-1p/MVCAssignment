package models

import (
	"errors"
	"inorder/pkg/types"
	"inorder/pkg/utils"
	"time"
)

func GetOrderedItems(orderID types.OrderID) ([]*types.OrderItem, error) {
	rows, err := db.Query("SELECT * FROM order_items WHERE order_id = ?", orderID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var otpt []*types.OrderItem = make([]*types.OrderItem, 0)
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
	var curr types.Order
	var rd types.MYSQLOrder

	err := db.QueryRow("SELECT * FROM orders WHERE id = ?", orderID).Scan(&curr.ID, &curr.IssuedBy, &rd.IssuedAt, &curr.Status, &rd.BillableAmount, &curr.TableNo, &rd.Waiter, &rd.PaidAt, &rd.Tip)
	if err != nil {
		return nil, err
	}
	curr.IssuedAt, err = time.Parse(time.DateTime, string(rd.IssuedAt))
	if err != nil {
		return nil, err
	}
	if len(rd.PaidAt) != 0 {
		curr.PaidAt, err = time.Parse(time.DateTime, string(rd.PaidAt))
		if err != nil {
			return nil, err
		}
	}
	if rd.BillableAmount.Valid {
		curr.BillableAmount = float32(rd.BillableAmount.Float64)
	}
	if rd.Waiter.Valid {
		curr.Waiter = types.UserID(rd.Waiter.Int64)
	}
	if rd.Tip.Valid {
		curr.Tip = float32(rd.Tip.Float64)
	}

	return &curr, nil
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
	order.IssuedAt = time.Now()
	order.Status = types.OrderStatusPending
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
	if err != nil {
		return []*types.Order{}, err
	}
	return utils.ParseOrderRows(rows)
}

func PayBill(order *types.Order, waiter types.UserID, tip float32) error {
	if order.Status != types.OrderStatusBilled {
		return errors.New("Order is not in billed state")
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
	Status         *types.OrderStatus `json:"status"`
	TableNo        *types.TableID     `json:"table_no"`
	BillableAmount *float32           `json:"billable_amount"`
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

func OrderNewItem(order *types.Order, itemID types.ItemID, quantity int, instructions string) (types.OrderItemID, error) {
	if quantity < 1 {
		return -1, errors.New("Quantity must be greater than or equal to 1")
	}
	if order.Status == types.OrderStatusBilled || order.Status == types.OrderStatusPaid {
		return -1, errors.New("Order status not open to item orders")
	}

	item, err := GetItemByID(itemID)
	if err != nil {
		return -1, err
	}

	var price float32 = float32(item.Price)

	res, err := db.Exec("INSERT INTO order_items (order_id, item_id, quantity, price,instructions) VALUES (?, ?, ?, ?, ?)", order.ID, item.ID, quantity, price, instructions)
	if err != nil {
		return -1, err
	}
	err = UpdateOrder(order, &OrderUpdateInstruction{
		Status: &types.OrderStatusPending,
	})
	if err != nil {
		return -1, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return -1, err
	}

	return types.OrderItemID(id), nil
}

func GetAllOrdersByUser(user *types.User, pg *types.Page) ([]*types.Order, error) {
	rows, err := db.Query("SELECT * FROM orders WHERE issued_by = ? LIMIT ? OFFSET ?", user.ID, pg.Limit, pg.Offset)
	if err != nil {
		return []*types.Order{}, err
	}
	return utils.ParseOrderRows(rows)
}
