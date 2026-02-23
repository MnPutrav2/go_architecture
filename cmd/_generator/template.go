package generate

import (
	"fmt"

	"github.com/MnPutrav2/go_architecture/cmd/_generator/temp"
)

func Tempate(name, ty string) {
	switch ty {
	case "-s":
		temp.TemplateService(name)
		fmt.Println("Generating success")

	case "-r":
		temp.TemplateRepo(name)
		fmt.Println("Generating success")

	case "-h":
		temp.TemplateHandle(name)
		fmt.Println("Generating success")

	case "-a":
		temp.TemplateRepo(name)
		temp.TemplateService(name)
		temp.TemplateHandle(name)
		fmt.Println("Generating success")

	default:
		fmt.Println("command not found")
	}
}
