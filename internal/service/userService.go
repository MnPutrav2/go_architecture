package service

import (
	"github.com/MnPutrav2/go_architecture/internal/repository"
)

type UserService struct {
	repo repository.UserRepository
}

func InitUserService(repo repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

// Entry
	