package types

type ItemID int

type Item struct {
	ID          ItemID
	Name        string
	Description string
	Price       float64
	Image       string
	Tags        []TagName
}
