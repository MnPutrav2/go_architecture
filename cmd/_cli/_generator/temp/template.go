package temp

import (
	"fmt"
)

func TemplateHandle(name string) {
	temp := fmt.Sprintf(`package handler

import (
	%sService "%s/internal/service/%s"

)

type %sHandle struct {
	service %sService.%sService
}

func Init%sHandle(service %sService.%sService) *%sHandle {
	return &%sHandle{service: service}
}

// Entry
`, name, moduleReader(), name, capitalize(name), name, capitalize(name), capitalize(name), name, capitalize(name), capitalize(name), capitalize(name))

	route := fmt.Sprintf(`
package route

import (
	"%s/internal/http/handler"
	%sService "%s/internal/service/%s"
	jwtEnc "%s/pkg/auth/jwt"
	"%s/pkg/middleware"
	"%s/pkg/response"
	"net/http"
)

func %sRoute(service %sService.%sService) http.HandlerFunc {
	h := handler.Init%sHandle(service)

	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			// write here
		default:
			response.Message("method not allowed", "method not allowed", "WARN", http.StatusMethodNotAllowed, w, r)
		}
	}
}
	`, moduleReader(), name, moduleReader(), name, moduleReader(), moduleReader(), moduleReader(), capitalize(name), name, capitalize(name), capitalize(name))

	handle := process2(temp, "http/handler", name)
	routes := process2(route, "http/route", name)
	fmt.Println(handle)
	fmt.Println(routes)
}

func TemplateRepo(name string) {
	temp := fmt.Sprintf(`package %sRepository

import (
	"database/sql"
)

type %sRepository struct {
	db *sql.DB
}

type %sRepository interface {
	// write in here
}

func Init%sRepository(db *sql.DB) %sRepository {
	return &%sRepository{db: db}
}

// Entry
	`, name, name, capitalize(name), capitalize(name), capitalize(name), name)

	tempq := fmt.Sprintf(`package %sRepository

var (
	// query = 
)
	`, name)

	handle := process(temp, "repository/", name, "impl")
	handle2 := process(tempq, "repository/", name, "query")
	fmt.Println(handle)
	fmt.Println(handle2)
}

func TemplateService(name string) {
	temp := fmt.Sprintf(`package %sService

import (
	%sRepository "%s/internal/repository/%s"
)

type %sService struct {
	repo %sRepository.%sRepository
}

type %sService interface {
	// write in here
}

func Init%sService(repo %sRepository.%sRepository) %sService {
	return &%sService{repo: repo}
}

// Entry
	`, name, name, moduleReader(), name, name, name, capitalize(name), capitalize(name), capitalize(name), name, capitalize(name), capitalize(name), name)

	handle := process(temp, "service/", name, "impl")
	fmt.Println(handle)
}
