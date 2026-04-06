package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/MnPutrav2/go_architecture/pkg/response"
)

func CTJson(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ct := r.Header.Get("Content-Type")
		if !strings.HasPrefix(ct, "application/json") {
			response.BadRequest("invalid content type, must be application/json", fmt.Errorf("invalid content type, must be application/json"), w, r)
			return
		}

		next.ServeHTTP(w, r)
	}
}

func CTFormData(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ct := r.Header.Get("Content-Type")
		if !strings.HasPrefix(ct, "multipart/form-data") {
			response.BadRequest("invalid content type, must be multipart/form-data", fmt.Errorf("invalid content type, must be multipart/form-data"), w, r)
			return
		}

		next.ServeHTTP(w, r)
	}
}
