package middleware

import (
	"context"
	"inorder/pkg/models"
	"inorder/pkg/types"
	"inorder/pkg/utils"
	"net/http"
)

func AuthenticationMiddleware(RefreshAuthToken bool) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			var err error
			var authToken utils.JSONWebToken

			authToken, err = utils.ExtractAuthToken(r)
			if err != nil {
				http.Error(w, "Unauthorized: "+err.Error(), http.StatusUnauthorized)
				return
			}

			err, claim := utils.VerifyJWT(authToken)
			if err != nil {
				http.Error(w, "Unauthorized: "+err.Error(), http.StatusUnauthorized)
				return
			}

			if !RefreshAuthToken && claim.Expired {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			userID := types.UserID(claim.Content["userID"].(float64))

			user, err := models.GetUserByID(userID)
			if err != nil {
				if err == utils.ErrUserNotFound {
					http.Error(w, utils.ErrUserNotFound.Error(), http.StatusNotFound)
					return
				}
				http.Error(w, err.Error(), http.StatusUnauthorized)
				return
			}
			ctx := context.WithValue(r.Context(), "user", user)

			if RefreshAuthToken {
				refreshToken, err := utils.ExtractRefreshToken(r)
				if err != nil {
					http.Error(w, err.Error(), http.StatusUnauthorized)
					return
				}
				err, refreshClaim := models.VerifyRefreshToken(refreshToken, user, true)
				if err != nil || refreshClaim.Expired {
					http.Error(w, "Unauthorized", http.StatusUnauthorized)
					return
				}
				newAuthToken, err := utils.CreateAuthToken(user)
				if err != nil {
					http.Error(w, "Internal Server Error: "+err.Error(), http.StatusInternalServerError)
					return
				}
				newRefreshToken, err := models.CreateRefreshToken(user)
				if err != nil {
					http.Error(w, "Internal Server Error: "+err.Error(), http.StatusInternalServerError)
					return
				}

				http.SetCookie(w, &http.Cookie{
					Name:     "authToken",
					Value:    string(newAuthToken),
					HttpOnly: true,
					Secure:   true,
					Path:     "/",
				})
				http.SetCookie(w, &http.Cookie{
					Name:     "refreshToken",
					Value:    string(newRefreshToken),
					HttpOnly: true,
					Secure:   true,
					Path:     "/api/auth",
				})
				ctx = context.WithValue(ctx, "authToken", newAuthToken)
				ctx = context.WithValue(ctx, "refreshToken", newRefreshToken)
			} else {
				ctx = context.WithValue(ctx, "authToken", authToken)
			}
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
