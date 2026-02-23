package cmd

import (
	"fmt"
	"net/http"

	"github.com/MnPutrav2/go_architecture/config"
	"github.com/MnPutrav2/go_architecture/internal/http/route"
	userRepository "github.com/MnPutrav2/go_architecture/internal/repository/user"
	userService "github.com/MnPutrav2/go_architecture/internal/service/user"
)

func Server(listen string) {

	db := config.InitDB()
	defer db.Close()

	mux := http.NewServeMux()

	// <----- Entry ----->

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	userRepo := userRepository.InitUserRepository(db)
	userServ := userService.InitUserService(userRepo)
	mux.HandleFunc("/user", route.UserRoute(userServ))

	// <----- Last ----->

	srv := &http.Server{
		Addr:    listen,
		Handler: mux,
	}

	fmt.Println("Server listen in port " + listen)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		panic(err)
	}
}
