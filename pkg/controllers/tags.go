package controllers

import (
	"encoding/json"
	"inorder/pkg/models"
	"net/http"
)

func GetAllTagsController(w http.ResponseWriter, r *http.Request) {
	tags, err := models.GetAllTags()
	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		http.Error(w, "Internal Server Error: "+err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(tags)
}
