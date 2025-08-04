package types

import "time"

type OrderID int

type Order struct {
	ID       OrderID
	IssuedBy UserID
	IssuedAt time.Time
	Status   OrderStatus
}
