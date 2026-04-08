package route

import (
	"database/sql"
	"net/http"

	"github.com/MnPutrav2/go_architecture/internal/http/handler"
	m "github.com/MnPutrav2/go_architecture/internal/http/route/method"
	"github.com/MnPutrav2/go_architecture/internal/repository"
	"github.com/MnPutrav2/go_architecture/internal/service"
)

func Route(db *sql.DB) *http.ServeMux {
	mux := http.NewServeMux()

	// [ Register route in here ]

	m.GET(mux, "/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	m.POST(mux, "/register", handler.CreateUserHandler(*service.InitUserService(*repository.InituserRepository(db))))

	// [ Register route in here ]

	return mux
}
