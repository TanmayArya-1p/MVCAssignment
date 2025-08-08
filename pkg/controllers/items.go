package controllers

import (
	"bytes"
	"encoding/json"
	"inorder/pkg/models"
	"inorder/pkg/types"
	"inorder/pkg/utils"
	"io"
	"net/http"
	"os"
	"slices"
	"strconv"

	"github.com/gorilla/mux"
)

func GetAllItemsController(w http.ResponseWriter, r *http.Request) {
	pg, err := utils.Paginate(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	items, err := models.GetAllItems(pg)
	if err != nil {
		http.Error(w, "Internal Server Error: "+err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(items)
}

func GetItemByIDController(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["itemid"]

	if id == "" {
		http.Error(w, "Missing id parameter", http.StatusBadRequest)
		return
	}
	tempID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid id parameter", http.StatusBadRequest)
		return
	}

	item, err := models.GetItemByID(types.ItemID(tempID))
	if err != nil {
		if err == utils.ErrItemNotFound {
			http.Error(w, "Item not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Internal Server Error: "+err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(item)
}

func CreateItemController(w http.ResponseWriter, r *http.Request) {

	var body ItemCRUDRequest
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	tags := []types.TagName{}
	for _, tag := range body.Tags {
		tags = append(tags, types.TagName(tag))
	}

	if body.Name == "" {
		http.Error(w, "Invalid name parameter", http.StatusBadRequest)
		return
	}
	if body.Price < 0 {
		http.Error(w, "Invalid price parameter", http.StatusBadRequest)
		return
	}

	var item types.Item
	item.Name = body.Name
	item.Description = body.Description
	item.Price = body.Price
	item.Tags = tags
	item.Image = body.Image

	_, err = models.CreateItem(&item)
	if err != nil {
		http.Error(w, "Internal Server Error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(item)
}

func DeleteItemController(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["itemid"]

	if id == "" {
		http.Error(w, "Missing id parameter", http.StatusBadRequest)
		return
	}
	tempID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid id parameter", http.StatusBadRequest)
		return
	}

	item, err := models.GetItemByID(types.ItemID(tempID))
	if err != nil {
		if err == utils.ErrItemNotFound {
			http.Error(w, "Item not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Internal Server Error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	err = models.DeleteItem(item)
	if err != nil {
		http.Error(w, "Internal Server Error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]any{"message": "Item deleted", "item": item})
}

func UpdateItemController(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["itemid"]

	if id == "" {
		http.Error(w, "Missing id parameter", http.StatusBadRequest)
		return
	}
	tempID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid id parameter", http.StatusBadRequest)
		return
	}

	var body ItemCRUDRequest
	err = json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if body.Price < 0 {
		http.Error(w, "Invalid price parameter", http.StatusBadRequest)
		return
	}
	var itemID types.ItemID = types.ItemID(tempID)

	item, err := models.GetItemByID(itemID)
	if err != nil {
		if err == utils.ErrItemNotFound {
			http.Error(w, "Item not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Internal Server Error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	var upd models.UpdateItemInstruction
	upd.Name = body.Name
	upd.Description = body.Description
	upd.Price = body.Price

	for _, tag := range body.Tags {
		upd.Tags = append(upd.Tags, types.TagName(tag))
	}

	upd.Image = body.Image

	err = models.UpdateItem(item, &upd)
	if err != nil {
		http.Error(w, "Internal Server Error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if slices.Equal(item.Tags, []types.TagName{"NOAC"}) {
		item.Tags = []types.TagName{}
	}

	json.NewEncoder(w).Encode(item)
}

func UploadImageController(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(32 << 20)
	var buf bytes.Buffer
	file, header, err := r.FormFile("image")
	if err != nil {
		http.Error(w, "Internal Server Error: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()
	io.Copy(&buf, file)

	filePath, fileURL, err := utils.GenerateImageUploadPath(header.Filename)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = os.WriteFile(filePath, buf.Bytes(), 0644)
	buf.Reset()

	json.NewEncoder(w).Encode(map[string]string{
		"message": "Uploaded successfully",
		"url":     fileURL,
	})
	return
}
