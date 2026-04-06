package temp

import (
	"fmt"
)

func TemplateHandle(name string) {
	temp := fmt.Sprintf(`package handler

import (
	"%s/internal/service"
	"net/http"
)

func RenameThisHandler(service service.%sService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		
		// 

	}
}

// Entry
`, moduleReader(), capitalize(name))

	handle := process2(temp, "http/handler", name+"Handler")
	fmt.Println(handle)
}

func TemplateRepo(name string) {
	temp := fmt.Sprintf(`package repository

import (
	"database/sql"
)

type %sRepository struct {
	db *sql.DB
}

func Init%sRepository(db *sql.DB) *%sRepository {
	return &%sRepository{db: db}
}

// Entry
	`, capitalize(name), name, capitalize(name), capitalize(name))

	handle := process2(temp, "repository/", name+"Repository")
	fmt.Println(handle)
}

func TemplateService(name string) {
	temp := fmt.Sprintf(`package service

import (
	"%s/internal/repository"
)

type %sService struct {
	repo repository.%sRepository
}

func Init%sService(repo repository.%sRepository) *%sService {
	return &%sService{repo: repo}
}

// Entry
	`, moduleReader(), capitalize(name), capitalize(name), capitalize(name), capitalize(name), capitalize(name), capitalize(name))

	handle := process2(temp, "service/", name+"Service")
	fmt.Println(handle)
}
