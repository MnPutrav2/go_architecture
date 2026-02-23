package generate

import (
	"fmt"

	"github.com/MnPutrav2/go_architecture/cmd/_generator/temp"
)

func Tempate(name, ty string) {
	switch ty {
	case "service":
		temp.TemplateService(name)
		fmt.Println("Generating success")

	case "repo":
		temp.TemplateRepo(name)
		fmt.Println("Generating success")

	case "handler":
		temp.TemplateHandle(name)
		fmt.Println("Generating success")

	case "all":
		temp.TemplateRepo(name)
		temp.TemplateService(name)
		temp.TemplateHandle(name)
		fmt.Println("Generating success")

	default:
		fmt.Println("command not found")
	}
}
