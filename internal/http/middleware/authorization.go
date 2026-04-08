package middleware

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	jwtEnc "github.com/MnPutrav2/go_architecture/pkg/auth/jwt"
	"github.com/MnPutrav2/go_architecture/pkg/response"
	"github.com/MnPutrav2/go_architecture/util"

	"github.com/joho/godotenv"
)

var _ = godotenv.Load()
var jwtKey = []byte(os.Getenv("JWT_SECURE_KEY"))

func Authorization(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		split := strings.SplitN(auth, " ", 2)

		if split[0] != "Bearer" {
			response.Unauthorization("Unauthorization", fmt.Errorf("Unauthorization"), w, r)
			return
		}

		claim, err := jwtEnc.ValidateJWT(split[1])
		if err != nil {
			response.Unauthorization("Unauthorization", err, w, r)
			return
		}

		ctx := context.WithValue(r.Context(), util.UserClaimsKey, claim)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}
