package repository

import (
	"context"
	"database/sql"

	"github.com/MnPutrav2/go_architecture/app/model"
	"github.com/MnPutrav2/go_architecture/app/pkg/query"
)

type AuthRepository struct {
	db *sql.DB
}

func InitauthRepository(db *sql.DB) *AuthRepository {
	return &AuthRepository{db: db}
}

// Entry
func (q *AuthRepository) InsertRefreshToken(ctx context.Context, token model.RefreshTokenPayload) error {
	if err := query.Init[model.RefreshToken](q.db).Insert(token).Exec(ctx); err != nil {
		return err
	}

	return nil
}

func (q *AuthRepository) CheckRefreshToken(ctx context.Context, token string) (model.RefreshToken, error) {
	result, err := query.Init[model.RefreshToken](q.db).Select("id, token_hash, expired").Find(ctx)
	if err != nil {
		return model.RefreshToken{}, err
	}

	return result, nil
}
