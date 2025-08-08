package utils

import "errors"

var (
	ErrUserNotFound  = errors.New("user not found")
	ErrInvalidRole   = errors.New("invalid role")
	ErrInvalidToken  = errors.New("invalid auth/refresh token")
	ErrInvalidJTI    = errors.New("invalid jti")
	ErrExpiredJTI    = errors.New("expired jti")
	ErrTagNotFound   = errors.New("tag not found")
	ErrItemNotFound  = errors.New("item not found")
	ErrOrderNotFound = errors.New("order not found")
	ErrInvalidLimit  = errors.New("invalid limit parameter")
	ErrInvalidOffset = errors.New("invalid offset parameter")
)
