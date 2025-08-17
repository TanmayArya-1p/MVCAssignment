package middleware

import (
	"context"
	"inorder/pkg/controllers"
	"inorder/pkg/models"
	"inorder/pkg/types"
	"inorder/pkg/utils"
	"net/http"
	"time"
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

			if val := authTokenCache.Get(authToken); val != nil {
				ctx := context.WithValue(r.Context(), types.AuthTokenContextKey, authToken)
				usr := val.(types.User)
				ctx = context.WithValue(r.Context(), types.UserContextKey, &usr)
				next.ServeHTTP(w, r.WithContext(ctx))
				return
			}

			err, claim := utils.VerifyJWT(authToken)
			if err != nil {
				http.Error(w, "Unauthorized: "+err.Error(), http.StatusUnauthorized)
				return
			}

			if !RefreshAuthToken && claim.Expired {
				http.Error(w, "Unauthorized: Token expired", http.StatusUnauthorized)
				return
			}
			userID := types.UserID(claim.Content["userID"].(float64))

			user, err := models.GetUserByID(userID)
			if err != nil {
				http.Error(w, err.Error(), http.StatusUnauthorized)
				return
			}
			ctx := context.WithValue(r.Context(), types.UserContextKey, user)

			if RefreshAuthToken {
				refreshToken, err := utils.ExtractRefreshToken(r)
				if err != nil {
					http.Error(w, err.Error(), http.StatusUnauthorized)
					return
				}
				err, refreshClaim := controllers.VerifyRefreshToken(refreshToken, user, true)
				if err != nil || refreshClaim.Expired {
					http.Error(w, "Unauthorized", http.StatusUnauthorized)
					return
				}
				newAuthToken, err := utils.CreateAuthToken(user)
				if err != nil {
					http.Error(w, "Internal Server Error: "+err.Error(), http.StatusInternalServerError)
					return
				}
				newRefreshToken, err := controllers.CreateRefreshToken(user)
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
				ctx = context.WithValue(ctx, types.AuthTokenContextKey, newAuthToken)
				ctx = context.WithValue(ctx, types.RefreshTokenContextKey, newRefreshToken)
			} else {
				ctx = context.WithValue(ctx, types.AuthTokenContextKey, authToken)
			}

			expClaim := claim.Content["exp"].(float64)
			expUnix := time.Unix(int64(expClaim), 0)
			ttl := time.Until(expUnix)
			authTokenCache.Set(authToken, *user, ttl)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
