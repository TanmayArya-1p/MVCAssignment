package controllers

import (
	"inorder/pkg/models"
	"inorder/pkg/types"
)

func ResolveBillableAmount(order *types.Order, toBill bool) ([]*types.OrderItem, error) {
	items, err := models.GetOrderedItems(order.ID)
	if err != nil {
		return []*types.OrderItem{}, err
	}

	var amt float32
	for _, item := range items {
		amt += item.Price * float32(item.Quantity)
	}

	var changeStatus *types.OrderStatus = nil
	if toBill {
		changeStatus = &types.OrderStatusBilled
	}

	err = models.UpdateOrder(order, &models.OrderUpdateInstruction{
		Status:         changeStatus,
		BillableAmount: &amt,
	})
	return items, err
}
