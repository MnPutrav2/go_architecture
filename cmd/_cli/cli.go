package main

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"syscall"

	"github.com/MnPutrav2/go_architecture/app/config"
	"github.com/MnPutrav2/go_architecture/app/migration"
	generate "github.com/MnPutrav2/go_architecture/cmd/_cli/_generator"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	db := config.InitDB()
	defer db.Close()

	if len(os.Args) < 2 {
		Help()
		return
	}

	switch os.Args[1] {
	case "install":
		if err := exec.Command("go", "mod", "download").Run(); err != nil {
			fmt.Printf("failed install depedency : %s", err)
			return
		}

		if err := exec.Command("npm", "install").Run(); err != nil {
			fmt.Printf("failed install depedency : %s", err)
			return
		}

		if err := exec.Command("cp", ".env.example", ".env"); err != nil {
			fmt.Printf("failed copy .env : %s", err)
			return
		}

		fmt.Println("Done.")
	case "migrate":
		migration.Auto(db)
	case "rollback":
		migration.Rollback(db)
	case "build":
		if err := os.MkdirAll("./build/app", 0755); err != nil {
			fmt.Printf("failed create folder : %s", err)
			return
		}
		fmt.Println("create output folder success : ./build/app/")

		if err := os.MkdirAll("./build/web", 0755); err != nil {
			fmt.Printf("failed create folder : %s", err)
			return
		}
		fmt.Println("create output folder success : ./build/web/")

		fmt.Println("try build app")
		if err := exec.Command("go", "build", "-o", "./build/app/server", "./cmd/server").Run(); err != nil {
			fmt.Printf("build failed : %s", err)
			return
		}

		fmt.Println("success build app")
		fmt.Println("try build web")
		if err := exec.Command("npm", "run", "build").Run(); err != nil {
			fmt.Printf("build web failed : %s", err)
			return
		}

		fmt.Println("success build web")
	case "help":
		Help()
		return
	case "dev":
		runDev()
	case "make:template":
		_, name, _ := strings.Cut(os.Args[2], "=")
		_, ty, _ := strings.Cut(os.Args[3], "=")

		generate.Tempate(name, ty)
	default:
		fmt.Println("command not found, use 'make help' for see available commands")
	}

}

func runDev() {
	fmt.Println("Starting development server...")

	backend := exec.Command("go", "run", "./cmd/server")

	backend.Stdout = os.Stdout
	backend.Stderr = os.Stderr

	frontend := exec.Command("npm", "run", "dev")
	frontend.Dir = "./"

	frontend.Stdout = os.Stdout
	frontend.Stderr = os.Stderr

	if err := backend.Start(); err != nil {
		fmt.Println("Failed to start backend:", err)
		return
	}

	fmt.Println("Backend started")

	if err := frontend.Start(); err != nil {
		fmt.Println("Failed to start frontend:", err)

		backend.Process.Kill()
		return
	}

	fmt.Println("Frontend started")
	fmt.Println()
	fmt.Println("Development server running")
	fmt.Println("Backend  : http://localhost:8080")
	fmt.Println("Frontend : http://localhost:5173")
	fmt.Println()
	fmt.Println("Press Ctrl+C to stop")

	signalChan := make(chan os.Signal, 1)

	signal.Notify(
		signalChan,
		os.Interrupt,
		syscall.SIGTERM,
	)

	<-signalChan

	fmt.Println("\nStopping development server...")

	// Stop backend
	if backend.Process != nil {
		backend.Process.Kill()
	}

	// Stop frontend
	if frontend.Process != nil {
		frontend.Process.Kill()
	}

	fmt.Println("Development server stopped")
}

func Help() {
	fmt.Print(`Available commands:
		
make template name=<file_name> type=<type>		 create template, example : make template name=user type=all
make run 						 running server
make build 						 build project for production
	`)
}
