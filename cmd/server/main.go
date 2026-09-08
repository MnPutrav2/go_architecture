package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/MnPutrav2/go_architecture/app/config"
	route "github.com/MnPutrav2/go_architecture/routes"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	db := config.InitDB()
	defer db.Close()
	mux := http.NewServeMux()

	listen := os.Getenv("LISTEN_PROD")
	srv := &http.Server{
		Addr:    listen,
		Handler: route.Route(mux, db),
	}

	fmt.Println("--- [ APP LOG ] ---")
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		panic(err)
	}
}
