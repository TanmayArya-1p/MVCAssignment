package controllers

import "inorder/pkg/types"

type ItemCRUDRequest struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Price       float64  `json:"price"`
	Tags        []string `json:"tags"`
	Image       string   `json:"image"`
}

type UserCRUDRequest struct {
	Username string     `json:"username"`
	Password string     `json:"password"`
	Role     types.Role `json:"role"`
}
