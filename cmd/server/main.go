package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/MnPutrav2/go_architecture/config"
	"github.com/MnPutrav2/go_architecture/internal/http/route"
	"github.com/MnPutrav2/go_architecture/internal/migration"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	db := config.InitDB()
	defer db.Close()

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
		Handler: route.Route(db),
	}

	list := fmt.Sprintf(`
====== GOLANG ======
			
server running  	: %s
start in		: %s 

--- [ APP LOG ] ---`, listen, time.Now().Format("2006-01-02 15:04:05"))

	fmt.Println(list)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		panic(err)
	}
}
