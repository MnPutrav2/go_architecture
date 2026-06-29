package service

import (
	"context"
	"fmt"
	"time"

	"github.com/MnPutrav2/go_architecture/app/model"
	jwtEnc "github.com/MnPutrav2/go_architecture/app/pkg/auth/jwt"
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

func (s *UserService) LoginUserService(ctx context.Context, username model.LoginUser) (model.Token, error) {
	result, err := s.repo.GetUserRepository(ctx, username.Name)
	if err != nil {
		return model.Token{}, err
	}

	if !password.Check(username.Password, result.Password) {
		return model.Token{}, fmt.Errorf("invalid username or password")
	}

	token, exp, err := jwtEnc.GenerateJWT(jwtEnc.User{
		UserID:   result.ID,
		Username: result.Name,
		Role:     "admin",
		Exp:      time.Now().Add(5),
	})

	if err != nil {
		return model.Token{}, err
	}

	return model.Token{Token: token, Expired: exp}, nil
}
