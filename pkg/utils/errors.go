package utils

import "errors"

var (
	ErrUserNotFound = errors.New("user not found")
	ErrInvalidRole  = errors.New("invalid role")
	ErrInvalidToken = errors.New("invalid auth/refresh token")
	ErrInvalidJTI   = errors.New("invalid jti")
	ErrExpiredJTI   = errors.New("expired jti")
	ErrTagNotFound  = errors.New("tag not found")
)
