package types

type ItemID int

type Item struct {
	ID          ItemID    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	Image       string    `json:"image"`
	Tags        []TagName `json:"tags"`
}

type UpdateItemInstruction struct {
	Name        string
	Description string
	Price       float64
	Image       string
	Tags        []TagName
}
