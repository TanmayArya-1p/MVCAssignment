package types

type OrderStatus string

var (
	OrderStatusPending   OrderStatus = "pending"
	OrderStatusPreparing OrderStatus = "preparing"
	OrderStatusServed    OrderStatus = "served"
	OrderStatusBilled    OrderStatus = "billed"
)
