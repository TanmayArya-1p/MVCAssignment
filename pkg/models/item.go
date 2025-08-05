package models

import (
	"database/sql"
	"errors"
	"fmt"
	"inorder/pkg/types"
	"inorder/pkg/utils"
	"slices"
	"strings"
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
	if err != nil {
		return item, err
	}
	var temp int64
	temp, err = row.LastInsertId()
	if err != nil {
		return item, err
	}
	item.ID = types.ItemID(temp)

	for _, tagName := range item.Tags {
		exists, tag := TagExists(tagName)
		if !exists {
			tag, err = CreateTag(tagName)
			if err != nil {
				return item, err
			}
		}
		GiveItemTag(tag, item.ID)
	}
	return item, nil
}

func DeleteItem(item *types.Item) (*types.Item, error) {

	_, err := db.Exec("DELETE FROM order_items WHERE item_id = ?", item.ID)
	if err != nil {
		return item, err
	}
	_, err = db.Exec("DELETE FROM items WHERE id = ?", item.ID)
	if err != nil {
		return item, err
	}
	return item, nil
}

func GetItemByID(itemID types.ItemID) (*types.Item, error) {
	var row *sql.Row
	var err error

	var item types.Item

	row = db.QueryRow("SELECT (id,name, description, price, image)  FROM items WHERE id = ?", itemID)
	err = row.Scan(&item.ID, &item.Name, &item.Description, &item.Price, &item.Image)

	var tempTags []*types.Tag
	tempTags, err = GetAllItemTags(itemID)
	if err != nil {
		return &item, err
	}

	for _, tag := range tempTags {
		item.Tags = append(item.Tags, tag.Name)
	}
	return &item, err
}

func GetAllItems(page types.Page) ([]*types.Item, error) {
	var rows *sql.Rows
	var err error

	var otpt []*types.Item

	if page.Limit != types.DefaultLimit {
		rows, err = db.Query("SELECT (id,name, description, price, image) FROM items LIMIT ? OFFSET ?", page.Limit, page.Offset)
	} else {
		rows, err = db.Query("SELECT (id,name, description, price, image) FROM items OFFSET ?", page.Offset)
	}
	if err != nil {
		return otpt, err
	}

	if exists := rows.Next(); !exists {
		return otpt, nil
	}
	for {
		var curr types.Item
		rows.Scan(&curr.ID, &curr.Name, &curr.Description, &curr.Price, &curr.Image)

		var tempTags []*types.Tag
		tempTags, err = GetAllItemTags(curr.ID)
		if err != nil {
			return otpt, err
		}

		for _, tag := range tempTags {
			curr.Tags = append(curr.Tags, tag.Name)
		}

		otpt = append(otpt, &curr)

		if isNext := rows.Next(); !isNext {
			break
		}
	}
	return otpt, nil
}

func GetAllItemsOfTag(tags []types.TagName) ([]*types.Item, error) {

	var inPl []string
	for _ = range len(tags) {
		inPl = append(inPl, "?")
	}

	var otpt []*types.Item

	var prep string = strings.Join(inPl, ",")
	var queryString string = fmt.Sprintf("SELECT DISTINCT  items.id,items.name,items.description,items.price,items.image FROM items INNER JOIN tag_rel ON items.id=tag_rel.item_id LEFT JOIN tags ON tags.id=tag_rel.tag_id WHERE tags.name IN (%s)", prep)
	rows, err := db.Query(queryString, tags)
	if err != nil {
		return otpt, err
	}
	if exist := rows.Next(); !exist {
		return otpt, nil
	}

	for {
		var curr types.Item
		rows.Scan(&curr.ID, &curr.Name, &curr.Description, &curr.Price, &curr.Image)

		var tempTags []*types.Tag
		tempTags, err = GetAllItemTags(curr.ID)
		if err != nil {
			return otpt, err
		}

		for _, tag := range tempTags {
			curr.Tags = append(curr.Tags, tag.Name)
		}

		otpt = append(otpt, &curr)

		if isNext := rows.Next(); !isNext {
			break
		}
	}

	var andRes []*types.Item
	for _, item := range otpt {
		if utils.SubsetOf(item.Tags, tags) {
			andRes = append(andRes, item)
		}
	}

	return andRes, nil
}

type UpdateItemInstruction struct {
	Name        string
	Description string
	Price       float64
	Image       string
	Tags        []types.TagName
}

func UpdateItem(item *types.Item, upd *UpdateItemInstruction) error {
	var updTags bool = true
	if upd.Name != "" {
		item.Name = upd.Name
	}
	if upd.Description != "" {
		item.Description = upd.Description
	}
	if upd.Price != 0 {
		item.Price = upd.Price
	}
	if upd.Image != "" {
		item.Image = upd.Image
	}
	if slices.Equal(upd.Tags, []types.TagName{"NOAC"}) {
		updTags = false
		item.Tags = upd.Tags
	}

	var queryString string = fmt.Sprintf("UPDATE items SET name=?, description=?, price=?, image=? WHERE id=?")
	_, err := db.Exec(queryString, item.Name, item.Description, item.Price, item.Image, item.ID)
	if err != nil {
		return err
	}

	if updTags {
		diff := utils.DiffCalculate(item.Tags, upd.Tags)
		for _, rem := range diff.Removed {
			err = RemoveItemTagByName(item.ID, rem)
			if err != nil {
				return err
			}
		}
		for _, add := range diff.Added {
			err = GiveItemTagByName(add, item.ID)
			if err != nil {
				return err
			}
		}
		item.Tags = upd.Tags
	}
	return err
}
