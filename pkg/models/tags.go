package models

import (
	"errors"
	"inorder/pkg/types"
)

type TagID int

type Tag struct {
	ID   TagID
	Name string
}

func CreateTag(name string) (*Tag, error) {
	res, err := db.Exec("INSERT INTO tags (name) VALUES (?)", name)
	if err != nil {
		return nil, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	return &Tag{ID: TagID(id), Name: name}, nil
}

func GetTag(id TagID) (*Tag, error) {
	//cachable
	var tag Tag
	err := db.QueryRow("SELECT id, name FROM tags WHERE id = ?", id).Scan(&tag.ID, &tag.Name)
	if err != nil {
		return nil, err
	}
	return &tag, nil
}

func TagExists(name string) (bool, *Tag) {
	//cachable
	var tag Tag
	err := db.QueryRow("SELECT id,name FROM tags WHERE name = ?", name).Scan(&tag.ID, &tag.Name)
	if err != nil {
		return false, nil
	}
	return true, &tag
}

func (tag *Tag) DeleteTag() error {
	//need todelete cache if user calls
	_, err := db.Exec("DELETE FROM tags WHERE id = ?", tag.ID)
	if err != nil {
		return err
	}
	return nil
}

func DeleteTagByID(tagid TagID) error {
	//delete cache
	var temp Tag = Tag{ID: tagid}
	return temp.DeleteTag()
}

func (tag *Tag) GiveItemTag(itemID types.ItemID) error {
	_, err := db.Exec("INSERT INTO tag_rel (item_id, tag_id) VALUES (?, ?)", itemID, tag.ID)
	if err != nil {
		return err
	}
	return nil
}

func GiveItemTagByName(tag_name string, itemID types.ItemID) error {
	exists, tag := TagExists(tag_name)
	if !exists {
		return errors.New("Tag does notag_namet exist")
	}
	return tag.GiveItemTag(itemID)
}

func (tag *Tag) RemoveItemTag(itemID types.ItemID) error {
	_, err := db.Exec("DELETE FROM tag_rel WHERE item_id = ? AND tag_id = ?", itemID, tag.ID)
	if err != nil {
		return err
	}
	return nil
}

func RemoveItemTagByName(itemID types.ItemID, tagName string) error {
	exists, tag := TagExists(tagName)
	if !exists {
		return errors.New("Tag does not exist")
	}
	return tag.RemoveItemTag(itemID)
}

func (tag *Tag) DeleteAllItems() error {
	_, err := db.Exec("DELETE FROM tag_rel WHERE tag_id = ?", tag.ID)
	if err != nil {
		return err
	}
	return nil
}

func DeleteAllItemTags(itemID types.ItemID) error {
	_, err := db.Exec("DELETE FROM tag_rel WHERE item_id = ?", itemID)
	if err != nil {
		return err
	}
	return nil
}

func GetAllTags() ([]*Tag, error) {
	var tags []*Tag
	rows, err := db.Query("SELECT id,name FROM tags")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if nonempty := rows.Next(); !nonempty {
		return tags, nil
	}

	for {
		var tag Tag
		if err := rows.Scan(&tag.ID, &tag.Name); err != nil {
			return nil, err
		}
		tags = append(tags, &tag)
		next := rows.Next()
		if !next {
			break
		}
	}
	return tags, nil
}

func GetAllItemTags(itemID types.ItemID) ([]*Tag, error) {
	rows, err := db.Query("SELECT tag_id FROM tag_rel WHERE item_id = ?", itemID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var otpt []*Tag
	if exists := rows.Next(); !exists {
		return otpt, nil
	}
	for {
		var tagId TagID
		if err := rows.Scan(&tagId); err != nil {
			return nil, err
		}
		tag, err := GetTag(tagId)
		if err != nil {
			return nil, err
		}
		otpt = append(otpt, tag)
		if isNext := rows.Next(); !isNext {
			break
		}
	}
	return otpt, nil
}
