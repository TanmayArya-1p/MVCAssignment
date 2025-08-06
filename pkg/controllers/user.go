package controllers

import (
	"encoding/json"
	"inorder/pkg/models"
	"inorder/pkg/types"
	"inorder/pkg/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetAllUsersController(w http.ResponseWriter, r *http.Request) {
	limit := r.URL.Query().Get("limit")
	offset := r.URL.Query().Get("offset")

	var pg types.Page
	var err error

	if limit == "" {
		pg.Limit = 10
	} else {
		pg.Limit, err = strconv.Atoi(limit)
		if err != nil {
			http.Error(w, "Invalid limit parameter", http.StatusBadRequest)
			return
		}
	}

	if offset == "" {
		pg.Offset = types.DefaultOffset
	} else {
		pg.Offset, err = strconv.Atoi(offset)
		if err != nil {
			http.Error(w, "Invalid offset parameter", http.StatusBadRequest)
			return
		}
	}

	users, err := models.GetAllUsers(pg)
	if err != nil {
		http.Error(w, "Internal Server Error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(users)
}

func GetUserByIDController(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["userid"]
	if id == "" {
		http.Error(w, "Missing id parameter", http.StatusBadRequest)
		return
	}

	uid, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid id parameter", http.StatusBadRequest)
		return
	}

	user, err := models.GetUserByID(types.UserID(uid))
	if err != nil {
		if err == utils.ErrUserNotFound {
			http.Error(w, "Invalid User ID", http.StatusBadRequest)
		} else {
			http.Error(w, "Internal Server Error: "+err.Error(), http.StatusInternalServerError)
		}
		return
	}
	json.NewEncoder(w).Encode(user)
}

func CreateUserController(w http.ResponseWriter, r *http.Request) {
	var user types.User

	var body map[string]string = make(map[string]string)
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	username := body["username"]
	password := body["password"]
	bodyRole := body["role"]

	if username == "" || password == "" || bodyRole == "" {
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}

	role := types.Role(bodyRole)
	if role != types.AdminRole && role != types.ChefRole && role != types.UserRole {
		http.Error(w, utils.ErrInvalidRole.Error(), http.StatusBadRequest)
		return
	}

	user.Username = username
	user.Password = password
	user.Role = role

	_, err = models.CreateUser(&user)
	if err != nil {
		http.Error(w, "Internal Server Error: "+err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(user)
}

func DeleteUserController(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["userid"]
	if id == "" {
		http.Error(w, "Missing id parameter", http.StatusBadRequest)
		return
	}

	uidInt, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid id parameter", http.StatusBadRequest)
		return
	}
	uid := types.UserID(uidInt)

	err = models.DeleteUserById(uid)
	if err != nil {
		http.Error(w, "Internal Server Error: "+err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"message": "User deleted successfully"})
}

func UpdateUserController(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["userid"]
	if id == "" {
		http.Error(w, "Missing id parameter", http.StatusBadRequest)
		return
	}

	uidInt, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid id parameter", http.StatusBadRequest)
		return
	}
	uid := types.UserID(uidInt)

	var body map[string]string = make(map[string]string)
	err = json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	username := body["username"]
	password := body["password"]
	bodyRole := body["role"]

	if username == "" && password == "" && bodyRole == "" {
		http.Error(w, "Missing atleast 1 field to update", http.StatusBadRequest)
		return
	}

	role := types.Role(bodyRole)
	if role != types.AdminRole && role != types.UserRole && role != types.ChefRole {
		http.Error(w, utils.ErrInvalidRole.Error(), http.StatusBadRequest)
		return
	}

	var user *types.User

	user, err = models.GetUserByID(uid)
	if err != nil {
		if err == utils.ErrUserNotFound {
			http.Error(w, utils.ErrUserNotFound.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, "Internal Server Error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	user, err = models.UpdateUser(&models.UserUpdateInstruction{
		User:              user,
		Username:          username,
		PlaintextPassword: password,
		Role:              role,
	})
	if err != nil {
		http.Error(w, "Internal Server Error: "+err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(user)
}
