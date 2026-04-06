package middleware

import (
	"fmt"
	"net/http"

	"github.com/MnPutrav2/go_architecture/pkg/response"
)

func RoleAdmin(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		ctx, ok := UserClaims(r.Context())
		if !ok {
			response.Unauthorization("unauthorization", fmt.Errorf("unauthorization"), w, r)
			return
		}

		if ctx.Role != "admin" {
			response.Forbidden("Access requires administrator privileges.", fmt.Errorf("Access requires administrator privileges."), w, r)
			return
		}

		next.ServeHTTP(w, r)
	}
}
