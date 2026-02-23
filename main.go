package main

import (
	"fmt"
	"os"

	"github.com/MnPutrav2/go_architecture/cmd"
	generate "github.com/MnPutrav2/go_architecture/cmd/_generator"

	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	if len(os.Args) < 2 {
		Help()
		return
	}

	dev := os.Getenv("LISTEN_DEVS")
	prod := os.Getenv("LISTEN_PROD")

	cli := os.Args[1]
	args := os.Args[2:]

	switch cli {
	case "dev":
		cmd.Server(dev)

	case "server":
		cmd.Server(prod)

	case "make:template":
		if len(args) == 0 || len(args[1]) == 0 {
			fmt.Println("Usage go run . make:template <file_name> <type>")
			return
		}

		generate.Tempate(args[0], args[1])

	default:
		fmt.Print("Command not found")
	}
}

func Help() {
	fmt.Print(`
Available commands:

	Application:
		go run . server						run server
		go run . dev						run dev
		
	Create template:
		go run . make:template <file_name> <type>		 <type> = [-h = create handler, -r = create repository, -s = create service, -a = create all]
	`)
}
