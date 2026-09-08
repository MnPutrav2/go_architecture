package temp

import (
	"fmt"
	"os"
	"strings"
)

func TemplateHandle(name string) {
	temp := fmt.Sprintf(`package handler

import (
	"context"
	"net/http"
	"%s/app/helper"

	"%s/app/service"
)

func RenameThisHandler(service service.%sService) http.HandlerFunc {
	return helper.Handler(func(ctx context.Context, w http.ResponseWriter, r *http.Request) {

		// Write here

	})
}

// Entry
`, moduleReader(), moduleReader(), capitalize(name))

	handle := process2(temp, "http/handler", name+"_handler")
	fmt.Println(handle)
}

func TemplateModel(name string) {
	temp := fmt.Sprintf(`package model

// Entry

type %s struct {
	// Input here
}`, capitalize(name))

	handle := process2(temp, "model/", name+"_model")
	fmt.Println(handle)

	if err := registerModel(capitalize(name)); err != nil {
		panic(err)
	}
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

	handle := process2(temp, "repository/", name+"_repository")
	fmt.Println(handle)
}

func TemplateService(name string) {
	temp := fmt.Sprintf(`package service

import (
	"%s/app/repository"
)

type %sService struct {
	repo repository.%sRepository
}

func Init%sService(repo repository.%sRepository) *%sService {
	return &%sService{repo: repo}
}

// Entry
	`, moduleReader(), capitalize(name), capitalize(name), capitalize(name), capitalize(name), capitalize(name), capitalize(name))

	handle := process2(temp, "service/", name+"_service")
	fmt.Println(handle)
}

func registerModel(name string) error {
	path := "app/util/registry_model.go"

	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	content := string(data)

	model := fmt.Sprintf("\tmodel.%s{},", name)

	if strings.Contains(content, model) {
		return nil
	}

	marker := "\t// @digo:models"

	index := strings.Index(content, marker)
	if index == -1 {
		return fmt.Errorf("marker %q not found", marker)
	}

	content = content[:index] + model + "\n" + content[index:]
	return os.WriteFile(path, []byte(content), 0644)
}
