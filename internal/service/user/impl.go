package userService

import (
	"context"

	userModel "github.com/MnPutrav2/go_architecture/internal/model/user"
	userRepository "github.com/MnPutrav2/go_architecture/internal/repository/user"
)

type userService struct {
	repo userRepository.UserRepository
}

type UserService interface {
	CreateUser(ctx context.Context, user userModel.Create) error
}

func InitUserService(repo userRepository.UserRepository) UserService {
	return &userService{repo: repo}
}

// Entry

func (s *userService) CreateUser(ctx context.Context, user userModel.Create) error {
	if err := s.repo.CreateUser(ctx, user); err != nil {
		return err
	}

	return nil
}
