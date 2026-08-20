package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/MnPutrav2/go_architecture/app/config"
	"github.com/MnPutrav2/go_architecture/app/migration"
	route "github.com/MnPutrav2/go_architecture/routes"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	db := config.InitDB()
	defer db.Close()
	mux := http.NewServeMux()

	auto, err := strconv.ParseBool(os.Getenv("AUTO_MIGRATE"))
	if err != nil {
		fmt.Println("env AUTO_MIGRATE need boolean")
		return
	}

	if auto {
		migration.Auto(db)
	}

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
