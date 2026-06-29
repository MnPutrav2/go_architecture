package route

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/MnPutrav2/go_architecture/app/http/handler"
	m "github.com/MnPutrav2/go_architecture/app/http/method"
	"github.com/MnPutrav2/go_architecture/app/repository"
	"github.com/MnPutrav2/go_architecture/app/service"
)

func Route(mux *http.ServeMux, db *sql.DB) http.Handler {
	// [ Register route in here ]

	// Web
	mux.Handle("/assets/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Serving : ", r.URL.Path)

		http.StripPrefix(
			"/assets/",
			http.FileServer(http.Dir("./public/assets")),
		).ServeHTTP(w, r)
	}))

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Serving : index.html")
		http.ServeFile(w, r, "./public/index.html")
	})
	// Web

	// API
	m.GET(mux, "/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	m.POST(mux, "/register", handler.CreateUserHandler(*service.InitUserService(*repository.InituserRepository(db))))
	m.GET(mux, "/users", handler.GetUserHandler(*service.InitUserService(*repository.InituserRepository(db))))

	// [ Register route in here ]

	return mux
}
