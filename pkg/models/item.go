package models

import (
	"database/sql"
	"errors"
	"inorder/pkg/types"
)

func CreateItem(item *types.Item) (*types.Item, error) {
	if item.Name == "" || item.Description == "" || item.Price == 0 {
		return item, errors.New("Invalid Parameters")
	}

	var row sql.Result
	var err error
	if item.Image != "" {
		row, err = db.Exec("INSERT INTO items (name, description, price, image) VALUES (?, ?, ?, ?)",
			item.Name, item.Description, item.Price, item.Image)
	} else {
		row, err = db.Exec("INSERT INTO items (name, description, price) VALUES (?, ?, ?)",
			item.Name, item.Description, item.Price)
	}

}
