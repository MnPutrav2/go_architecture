package service

import (
	"context"

	"github.com/MnPutrav2/go_architecture/app/model"
	"github.com/MnPutrav2/go_architecture/app/pkg/password"
	"github.com/MnPutrav2/go_architecture/app/repository"
)

type UserService struct {
	repo repository.UserRepository
}

func InitUserService(repo repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

// Entry
func (s *UserService) CreateUserService(ctx context.Context, request model.CreateUser) error {

	pw, err := password.Hash(request.Password)
	if err != nil {
		return err
	}

	payload := model.CreateUser{
		Name:     request.Name,
		Password: pw,
		Email:    request.Email,
	}

	if err := s.repo.CreateUserRepository(ctx, payload); err != nil {
		return err
	}

	return nil
}
