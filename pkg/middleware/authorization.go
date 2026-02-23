package middleware

import (
	"net/http"
	"os"
	"strings"

	jwtEnc "github.com/MnPutrav2/go_architecture/pkg/auth/jwt"
	"github.com/MnPutrav2/go_architecture/pkg/response"

	"github.com/joho/godotenv"
)

var _ = godotenv.Load()
var jwtKey = []byte(os.Getenv("JWT_SECURE_KEY"))

func Token(r *http.Request) string {

	auth := r.Header.Get("Authorization")
	// Check Header
	split := strings.SplitN(auth, " ", 2)

	if len(split) != 2 || split[0] != "Bearer" {
		return ""
	}

	if split[0] != "Bearer" {
		return ""
	}

	return split[1]
}

func Authorization(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		split := strings.SplitN(auth, " ", 2)

		if split[0] != "Bearer" {
			response.Message("token need bearer", "token need bearer", "WARN", 401, w, r)
			return
		}

		if _, err := jwtEnc.ValidateJWT(split[1]); err != nil {
			response.Message("unauthorization", err.Error(), "WARN", 401, w, r)
			return
		}

		next(w, r)
	}
}
