package types

type OrderStatus string

var (
	OrderStatusPending   OrderStatus = "pending"
	OrderStatusPreparing OrderStatus = "preparing"
	OrderStatusServed    OrderStatus = "served"
	OrderStatusBilled    OrderStatus = "billed"
	OrderStatusPaid      OrderStatus = "paid"
)

type OrderItemStatus string

var (
	OrderItemStatusPending   OrderItemStatus = "pending"
	OrderItemStatusPreparing OrderItemStatus = "preparing"
	OrderItemStatusServed    OrderItemStatus = "served"
)
