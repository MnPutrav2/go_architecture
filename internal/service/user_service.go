package service

import (
	"context"

	"github.com/MnPutrav2/go_architecture/internal/model"
	"github.com/MnPutrav2/go_architecture/internal/repository"
)

type UserService struct {
	repo repository.UserRepository
}

func InitUserService(repo repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

// Entry
func (s *UserService) CreateUserService(ctx context.Context, request model.CreateUser) error {
	if err := s.repo.CreateUserRepository(ctx, request); err != nil {
		return err
	}

	return nil
}
