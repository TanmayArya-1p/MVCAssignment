package middleware

import (
	"inorder/pkg/types"
	"inorder/pkg/utils"
	"net/http"
)

func AuthorizationMiddleware(PrivilegesLowerBound types.Role) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			user := r.Context().Value(types.UserContextKey).(*types.User)

			if utils.RolePrivileges[user.Role] < utils.RolePrivileges[PrivilegesLowerBound] {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
