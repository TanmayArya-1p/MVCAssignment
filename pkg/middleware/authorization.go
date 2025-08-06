package middleware

import (
	"inorder/pkg/types"
	"inorder/pkg/utils"
	"net/http"
)

func AuthorizationMiddleware(next http.Handler, PrivsLowerBound types.Role) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value("user").(*types.User)

		if utils.RolePrivs[user.Role] < utils.RolePrivs[PrivsLowerBound] {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}
