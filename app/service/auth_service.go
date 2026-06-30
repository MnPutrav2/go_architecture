package service

import (
	"context"
	"fmt"
	"time"

	httperror "github.com/MnPutrav2/go_architecture/app/http/error"
	"github.com/MnPutrav2/go_architecture/app/model"
	"github.com/MnPutrav2/go_architecture/app/pkg/auth"
	jwtEnc "github.com/MnPutrav2/go_architecture/app/pkg/auth/jwt"
	"github.com/MnPutrav2/go_architecture/app/pkg/password"
	"github.com/MnPutrav2/go_architecture/app/repository"
)

type AuthService struct {
	repo  repository.AuthRepository
	repo2 repository.UserRepository
}

func InitAuthService(repo repository.AuthRepository, repo2 repository.UserRepository) *AuthService {
	return &AuthService{repo: repo, repo2: repo2}
}

// Entry

func (s *AuthService) LoginService(ctx context.Context, username model.LoginUser) (model.Token, error) {
	result, err := s.repo2.GetUserRepository(ctx, username.Name)
	if err != nil {
		return model.Token{}, err
	}

	if !password.Check(username.Password, result.Password) {
		return model.Token{}, fmt.Errorf("invalid username or password")
	}

	refreshToken, hash, exp, err := auth.GenerateRefreshToken()
	if err != nil {
		return model.Token{}, err
	}

	if err := s.repo.InsertRefreshToken(ctx, model.RefreshTokenPayload{
		UserID:    result.ID,
		TokenHash: hash,
		ExpiredAt: exp,
	}); err != nil {
		return model.Token{}, err
	}

	token, exp, err := jwtEnc.GenerateJWT(jwtEnc.User{
		UserID:   result.ID,
		Username: result.Name,
		Role:     result.Role,
		Exp:      time.Now().Add(5),
	})

	if err != nil {
		return model.Token{}, err
	}

	return model.Token{Token: token, RefreshToken: refreshToken, Expired: exp}, nil
}

func (s *AuthService) RefreshToken(ctx context.Context, token string) (model.Token, error) {
	hash := auth.HashToken(token)
	user, err := s.repo.CheckRefreshToken(ctx, hash)
	if err != nil {
		return model.Token{}, fmt.Errorf("UNAUTHORIZATION: %w", httperror.ErrIsUnauthorization)
	}

	refreshToken, hash, exp, err := auth.GenerateRefreshToken()
	if err != nil {
		return model.Token{}, err
	}

	result, err := s.repo2.GetUserByIdRepository(ctx, user.UserID)
	if err != nil {
		return model.Token{}, err
	}

	if err := s.repo.InsertRefreshToken(ctx, model.RefreshTokenPayload{
		UserID:    result.ID,
		TokenHash: hash,
		ExpiredAt: exp,
	}); err != nil {
		return model.Token{}, err
	}

	token, exp, err = jwtEnc.GenerateJWT(jwtEnc.User{
		UserID:   result.ID,
		Username: result.Name,
		Role:     result.Role,
		Exp:      time.Now().Add(5),
	})

	if err != nil {
		return model.Token{}, err
	}

	return model.Token{Token: token, RefreshToken: refreshToken, Expired: exp}, nil
}
