package utils

import (
	"errors"
	"inorder/pkg/types"
	"net/http"
	"strings"
)

type Priority uint

var RolePrivs map[types.Role]Priority = map[types.Role]Priority{
	types.AdminRole: Priority(3),
	types.ChefRole:  Priority(2),
	types.UserRole:  Priority(1),
}

func ExtractAuthToken(req *http.Request) (JSONWebToken, error) {
	headerAuth := req.Header.Get("Authorization")
	if headerAuth == "" {
		cookie, err := req.Cookie("authToken")
		if err != nil {
			return "", err
		}
		if cookie.Value != "" {
			token := strings.Split(cookie.Value, " ")[1]
			return JSONWebToken(token), nil
		} else {
			return "", errors.New("Invalid Auth Token")
		}
	}
	token := strings.Split(headerAuth, " ")[1]
	return JSONWebToken(token), nil
}

func ExtractRefreshToken(req *http.Request) (JSONWebToken, error) {
	headerAuth := req.Header.Get("refreshToken")
	if headerAuth == "" {
		cookie, err := req.Cookie("refreshToken")
		if err != nil {
			return "", err
		}
		if cookie.Value != "" {
			token := strings.Split(cookie.Value, " ")[1]
			return JSONWebToken(token), nil
		} else {
			return "", errors.New("Invalid Refresh Token")
		}
	}
	token := strings.Split(headerAuth, " ")[1]
	return JSONWebToken(token), nil
}
