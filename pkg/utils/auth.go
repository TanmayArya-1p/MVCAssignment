package utils

import (
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
			return "", ErrInvalidToken
		}
		if cookie.Value != "" {
			return JSONWebToken(cookie.Value), nil
		} else {
			return "", ErrInvalidToken
		}
	}
	tokens := strings.Split(headerAuth, " ")
	if len(tokens) != 2 {
		return "", ErrInvalidToken
	}
	return JSONWebToken(tokens[1]), nil
}

func ExtractRefreshToken(req *http.Request) (JSONWebToken, error) {
	headerAuth := req.Header.Get("refreshToken")
	if headerAuth == "" {
		cookie, err := req.Cookie("refreshToken")
		if err != nil {
			return "", ErrInvalidToken
		}
		if cookie.Value != "" {
			return JSONWebToken(cookie.Value), nil
		} else {
			return "", ErrInvalidToken
		}
	}
	tokens := strings.Split(headerAuth, " ")
	if len(tokens) != 2 {
		return "", ErrInvalidToken
	}
	return JSONWebToken(tokens[1]), nil
}
