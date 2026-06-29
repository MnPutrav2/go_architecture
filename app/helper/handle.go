package helper

import (
	"context"
	"net/http"
	"time"
)

func Handler(fn func(ctx context.Context, w http.ResponseWriter, r *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, close := context.WithTimeout(r.Context(), 5*time.Second)
		defer close()

		fn(ctx, w, r)
	}
}
