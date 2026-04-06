package route

import (
	"database/sql"
	"net/http"

	"github.com/MnPutrav2/go_architecture/internal/http/handler"
	"github.com/MnPutrav2/go_architecture/internal/repository"
	"github.com/MnPutrav2/go_architecture/internal/service"
)

func Route(db *sql.DB) *http.ServeMux {
	mux := http.NewServeMux()

	// Add route in here

	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	mux.HandleFunc("POST /register", handler.CreateUserHandler(*service.InitUserService(*repository.InituserRepository(db))))

	// Add route in here

	return mux
}
