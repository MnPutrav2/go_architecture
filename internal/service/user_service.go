package service

import (
	"context"

	"github.com/MnPutrav2/go_architecture/internal/model"
	"github.com/MnPutrav2/go_architecture/internal/repository"
	"github.com/google/uuid"
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

func (s *UserService) GetUserService(ctx context.Context) ([]model.Users, error) {
	result, err := s.repo.GetUserRepository(ctx)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *UserService) DeleteUserService(ctx context.Context, id uuid.UUID) error {
	if err := s.repo.DeleteUserRepository(ctx, id); err != nil {
		return err
	}

	return nil
}
