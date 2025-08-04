package types

import "time"

type UserID int

type User struct {
	ID             UserID
	Username       string
	HashedPassword string
	Password       string
	Role           Role
	CreatedAt      time.Time
}
