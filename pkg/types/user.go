package types

import "time"

type UserID int

type User struct {
	ID             UserID    `json:"id"`
	Username       string    `json:"username"`
	HashedPassword string    `json:"hashed_password"`
	Password       string    `json:"password"`
	Role           Role      `json:"role"`
	CreatedAt      time.Time `json:"created_at"`
}
