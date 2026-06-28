package route

import (
	"database/sql"
	"net/http"

	"github.com/MnPutrav2/go_architecture/app/http/handler"
	m "github.com/MnPutrav2/go_architecture/app/http/method"
	"github.com/MnPutrav2/go_architecture/app/http/middleware"
	"github.com/MnPutrav2/go_architecture/app/repository"
	"github.com/MnPutrav2/go_architecture/app/service"
)

func RouteApi(mux *http.ServeMux, db *sql.DB) http.Handler {
	// [ Register route in here ]

	m.GET(mux, "/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	m.POST(mux, "/register", handler.CreateUserHandler(*service.InitUserService(*repository.InituserRepository(db))))
	m.GET(mux, "/users", handler.GetUserHandler(*service.InitUserService(*repository.InituserRepository(db))))
	m.DELETE(mux, "/users/{id}", handler.DeleteUserHandler(*service.InitUserService(*repository.InituserRepository(db))))

	// [ Register route in here ]

	// [ Add middleware ]

	md := middleware.HandeCORS(mux)

	// [ Add middleware ]
	return md
}
