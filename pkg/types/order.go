package types

import (
	"database/sql"
	"time"
)

type OrderID int
type TableID int

type Order struct {
	ID             OrderID     `json:"id"`
	IssuedBy       UserID      `json:"issued_by"`
	IssuedAt       time.Time   `json:"issued_at"`
	Status         OrderStatus `json:"status"`
	BillableAmount float32     `json:"billable_amount"`
	TableNo        TableID     `json:"table_no"`
	Waiter         UserID      `json:"waiter"`
	PaidAt         time.Time   `json:"paid_at"`
	Tip            float32     `json:"tip"`
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
