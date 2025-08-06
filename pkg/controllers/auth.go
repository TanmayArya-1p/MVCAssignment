package controllers

import (
	"encoding/json"
	"inorder/pkg/models"
	"inorder/pkg/types"
	"inorder/pkg/utils"
	"log"
	"net/http"
)

func RegisterController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var body map[string]any
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", http.StatusBadRequest)
		return
	}
	username := body["username"].(string)
	password := body["password"].(string)

	if username == "" || password == "" {
		http.Error(w, "Invalid username or password", http.StatusBadRequest)
		return
	}

	var user types.User = types.User{
		Username: username,
		Password: password,
	}

	uid, err := models.CreateUser(&user)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error: "+err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]any{
		"message": "user registered successfully",
		"user_id": int(uid),
	})
}

func LoginController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var body map[string]any
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", http.StatusBadRequest)
		return
	}
	username := body["username"].(string)
	password := body["password"].(string)

	if username == "" || password == "" {
		http.Error(w, "Invalid username or password", http.StatusBadRequest)
		return
	}

	user, err := models.GetUserByUsername(username)
	if err != nil {
		http.Error(w, "Internal Server Error :"+err.Error(), http.StatusInternalServerError)
		return
	}

	stat, err := utils.VerifyUser(user, password)
	if !stat {
		http.Error(w, "Invalid username or password", http.StatusBadRequest)
		return
	}
	authToken, err := utils.CreateAuthToken(user)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error: "+err.Error(), http.StatusInternalServerError)
		return
	}
	refreshToken, err := models.CreateRefreshToken(user)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	strAuth := string(authToken)
	strRefresh := string(refreshToken)

	http.SetCookie(w, &http.Cookie{
		Name:     "authToken",
		Value:    strAuth,
		HttpOnly: true,
		Secure:   true,
	})
	http.SetCookie(w, &http.Cookie{
		Name:     "refreshToken",
		Value:    strRefresh,
		HttpOnly: true,
		Secure:   true,
		Path:     "/api/auth",
	})

	json.NewEncoder(w).Encode(map[string]any{
		"message":      "user logged in successfully",
		"authToken":    strAuth,
		"refreshToken": strRefresh,
	})
}

func LogoutController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	refreshToken, err := utils.ExtractRefreshToken(r)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	user := r.Context().Value("user").(*types.User)
	models.VerifyRefreshToken(refreshToken, user, true)

	http.SetCookie(w, &http.Cookie{
		Name:     "authToken",
		Value:    "",
		HttpOnly: true,
		Secure:   true,
	})
	http.SetCookie(w, &http.Cookie{
		Name:     "refreshToken",
		Value:    "",
		HttpOnly: true,
		Secure:   true,
		Path:     "/api/auth",
	})

	json.NewEncoder(w).Encode(map[string]any{
		"message": "user logged out successfully",
	})
}

func VerifyController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]any{
		"message": "verified",
	})
}

func RefreshController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]any{
		"message": "tokens refreshed",
	})
}
