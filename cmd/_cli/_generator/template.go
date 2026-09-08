package generate

import (
	"fmt"

	"github.com/MnPutrav2/go_architecture/cmd/_cli/_generator/temp"
)

func Tempate(name, ty string) {
	switch ty {
	case "service":
		temp.TemplateService(name)

	case "repo":
		temp.TemplateRepo(name)

	case "handler":
		temp.TemplateHandle(name)

	case "model":
		temp.TemplateModel(name)

	case "all":
		temp.TemplateRepo(name)
		temp.TemplateService(name)
		temp.TemplateHandle(name)
		temp.TemplateModel(name)

	default:
		fmt.Println("command not found")
	}
}
