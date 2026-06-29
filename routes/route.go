package route

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/MnPutrav2/go_architecture/app/http/handler"
	m "github.com/MnPutrav2/go_architecture/app/http/method"
	"github.com/MnPutrav2/go_architecture/app/http/middleware"
	"github.com/MnPutrav2/go_architecture/app/repository"
	"github.com/MnPutrav2/go_architecture/app/service"
)

func Route(mux *http.ServeMux, db *sql.DB) http.Handler {
	// [ Register route in here ]

	// Web
	mux.Handle("/assets/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Serving : ", r.URL.Path)
		http.StripPrefix("/assets/", http.FileServer(http.Dir("./final/web/assets"))).ServeHTTP(w, r)
	}))

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Serving : index.html")
		http.ServeFile(w, r, "./final/web/index.html")
	})
	// Web

	// API
	m.GET(mux, "/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	userService := service.InitUserService(*repository.InituserRepository(db))

	m.POST(mux, "/register", handler.CreateUserHandler(*userService))
	m.POST(mux, "/login", handler.LoginUserHandler(*service.InitAuthService(*repository.InitauthRepository(db), *repository.InituserRepository(db))))

	// [ Register route in here ]

	mode := os.Getenv("MODE")
	if mode == "PROD" {
		return mux
	} else {
		mx := middleware.HandeCORS(mux)
		return mx
	}

}
