package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	if len(os.Args) < 2 {
		return
	}

	if len(os.Args) < 4 {
		fmt.Println("Usage: make:template <name> <type>")
		return
	}

	switch os.Args[1] {
	case "make:template":
		_, name, _ := strings.Cut(os.Args[2], "=")
		_, ty, _ := strings.Cut(os.Args[3], "=")

		fmt.Println("Name:", name)
		fmt.Println("Type:", ty)
	default:
		fmt.Println("more")
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
