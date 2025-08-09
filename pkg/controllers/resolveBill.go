package controllers

import (
	"inorder/pkg/models"
	"inorder/pkg/types"
)

func ResolveBillableAmount(order *types.Order, toBill bool) error {
	items, err := models.GetOrderedItems(order.ID)
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

	return models.UpdateOrder(order, &models.OrderUpdateInstruction{
		Status:         changeStatus,
		BillableAmount: &amt,
	})
}
