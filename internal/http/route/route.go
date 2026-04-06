package route

import (
	"database/sql"
	"net/http"
)

func Route(db *sql.DB) *http.ServeMux {
	mux := http.NewServeMux()

	// Add route in here

	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	// Add route in here

	return mux
}
