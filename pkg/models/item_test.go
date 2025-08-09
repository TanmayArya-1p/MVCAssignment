package models

import (
	"inorder/pkg/types"
	"reflect"
	"testing"

	"github.com/google/uuid"
)

func TestItemsCRUD(t *testing.T) {
	//creating item
	item := types.Item{
		Name:        "item_" + uuid.New().String(),
		Description: "test desc",
		Price:       10.0,
	}
	_, err := CreateItem(&item)
	if err != nil {
		t.Errorf("Error creating item: %v", err)
	}
	//check if its in db
	it, err := GetItemByID(item.ID)
	if err != nil {
		t.Errorf("Error getting item: %v", err)
	}
	if reflect.DeepEqual(it, item) {
		t.Error("Item not found in db", it, item)
	}
	//delete and check if its gone
	err = DeleteItem(it)
	if err != nil {
		t.Errorf("Error deleting item: %v", err)
	}
	it, err = GetItemByID(item.ID)
	if err == nil {
		t.Errorf("Item still in db after deletion %v", it)
	}
	//again create item
	item = types.Item{
		Name:        "item_" + uuid.New().String(),
		Description: "test desc",
		Price:       10.0,
	}
	_, err = CreateItem(&item)
	if err != nil {
		t.Errorf("Error creating item: %v", err)
	}
	//update the item
	err = UpdateItem(&item, &types.UpdateItemInstruction{
		Name:        "item_" + uuid.New().String(),
		Description: "newdesc",
		Price:       20.0,
		Tags:        []types.TagName{"hot", "spicy"},
	})
	if err != nil {
		t.Errorf("Error updating item: %v", err)
	}
	//check if item updated in db
	it, err = GetItemByID(item.ID)
	if err != nil {
		t.Errorf("Error getting item: %v", err)
	}
	if reflect.DeepEqual(it, item) {
		t.Error("Item not updated in db", it, item)
	}
	DeleteItem(it)
	t.Log("Passed all item tests")

}
