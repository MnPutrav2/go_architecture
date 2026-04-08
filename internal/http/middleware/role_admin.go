package middleware

import (
	"fmt"
	"net/http"
	"strings"

	jwtEnc "github.com/MnPutrav2/go_architecture/pkg/auth/jwt"
	"github.com/MnPutrav2/go_architecture/pkg/response"
	"github.com/MnPutrav2/go_architecture/util"
)

func RoleAdmin(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		ctx, ok := r.Context().Value(util.UserClaimsKey).(*jwtEnc.Claims)
		if !ok {
			response.Unauthorization("unauthorization", fmt.Errorf("unauthorization"), w, r)
			return
		}

		if strings.ToLower(ctx.Role) != "admin" {
			response.Forbidden("Access requires administrator privileges.", fmt.Errorf("Access requires administrator privileges."), w, r)
			return
		}

		next.ServeHTTP(w, r)
	}
}
