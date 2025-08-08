package types

import (
	"database/sql"
	"time"
)

type OrderID int
type TableID int

type Order struct {
	ID             OrderID
	IssuedBy       UserID
	IssuedAt       time.Time
	Status         OrderStatus
	BillableAmount float32
	TableNo        TableID
	Waiter         UserID
	PaidAt         time.Time
	Tip            float32
}

type MYSQLOrder struct {
	IssuedAt       []byte
	BillableAmount sql.NullFloat64
	Waiter         sql.NullInt64
	PaidAt         []byte
	Tip            sql.NullFloat64
}

type OrderItemID int
type OrderItem struct {
	ID           OrderItemID
	OrderID      OrderID
	ItemID       ItemID
	Instructions string
	Quantity     int
	Price        float32
	IssuedAt     time.Time
	Status       OrderItemStatus
}
