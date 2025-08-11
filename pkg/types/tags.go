package types

type TagID int
type TagName string

type Tag struct {
	ID   TagID   `json:"id"`
	Name TagName `json:"name"`
}
