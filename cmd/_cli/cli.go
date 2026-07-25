package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

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
	case "make:template":
		_, name, _ := strings.Cut(os.Args[2], "=")
		_, ty, _ := strings.Cut(os.Args[3], "=")

		generate.Tempate(name, ty)
	default:
		fmt.Println("command not found, use 'make help' for see available commands")
	}

}

func Help() {
	fmt.Print(`Available commands:
		
make template name=<file_name> type=<type>		 create template, example : make template name=user type=all
make run 						 running server
make build 						 build project for production
	`)
}
