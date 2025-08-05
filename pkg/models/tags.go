package models

import (
	"errors"
	"inorder/pkg/types"
)

func CreateTag(name types.TagName) (*types.Tag, error) {
	res, err := db.Exec("INSERT INTO tags (name) VALUES (?)", name)
	if err != nil {
		return nil, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	return &types.Tag{ID: types.TagID(id), Name: name}, nil
}

func GetTag(id types.TagID) (*types.Tag, error) {
	//cachable
	var tag types.Tag
	err := db.QueryRow("SELECT id, name FROM tags WHERE id = ?", id).Scan(&tag.ID, &tag.Name)
	if err != nil {
		return nil, err
	}
	return &tag, nil
}

func TagExists(name types.TagName) (bool, *types.Tag) {
	//cachable
	var tag types.Tag
	err := db.QueryRow("SELECT id,name FROM tags WHERE name = ?", name).Scan(&tag.ID, &tag.Name)
	if err != nil {
		return false, nil
	}
	return true, &tag
}

func DeleteTag(tag *types.Tag) error {
	//need todelete cache if user calls
	_, err := db.Exec("DELETE FROM tags WHERE id = ?", tag.ID)
	if err != nil {
		return err
	}
	return nil
}

func DeleteTagByID(tagid types.TagID) error {
	//delete cache
	var temp types.Tag = types.Tag{ID: tagid}
	return DeleteTag(&temp)
}

func GiveItemTag(tag *types.Tag, itemID types.ItemID) error {
	_, err := db.Exec("INSERT INTO tag_rel (item_id, tag_id) VALUES (?, ?)", itemID, tag.ID)
	if err != nil {
		return err
	}
	return nil
}

func GiveItemTagByName(tag_name types.TagName, itemID types.ItemID) error {
	exists, tag := TagExists(tag_name)
	if !exists {
		return errors.New("Tag does not exist")
	}
	return GiveItemTag(tag, itemID)
}

func RemoveItemTag(tag *types.Tag, itemID types.ItemID) error {
	_, err := db.Exec("DELETE FROM tag_rel WHERE item_id = ? AND tag_id = ?", itemID, tag.ID)
	if err != nil {
		return err
	}
	return nil
}

func RemoveItemTagByName(itemID types.ItemID, tagName types.TagName) error {
	exists, tag := TagExists(tagName)
	if !exists {
		return errors.New("Tag does not exist")
	}
	return RemoveItemTag(tag, itemID)
}

func DeleteAllItems(tag *types.Tag) error {
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

func GetAllTags() ([]*types.Tag, error) {
	var tags []*types.Tag
	rows, err := db.Query("SELECT id,name FROM tags")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if nonempty := rows.Next(); !nonempty {
		return tags, nil
	}

	for {
		var tag types.Tag
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

func GetAllItemTags(itemID types.ItemID) ([]*types.Tag, error) {
	rows, err := db.Query("SELECT tag_id FROM tag_rel WHERE item_id = ?", itemID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var otpt []*types.Tag
	if exists := rows.Next(); !exists {
		return otpt, nil
	}
	for {
		var tagId types.TagID
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
