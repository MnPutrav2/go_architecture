package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/MnPutrav2/go_architecture/config"
	"github.com/MnPutrav2/go_architecture/internal/http/route"
)

func main() {
	db := config.InitDB()
	defer db.Close()

	listen := os.Getenv("LISTEN_PROD")
	srv := &http.Server{
		Addr:    listen,
		Handler: route.Route(db),
	}

	fmt.Println("Server listen in port " + listen)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		panic(err)
	}
}
