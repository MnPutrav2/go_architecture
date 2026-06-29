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
func (q *AuthRepository) InsertRefreshToken(ctx context.Context, token model.RefreshTokenInsert) error {
	if err := query.Init[model.RefreshToken](q.db).Insert(token).Exec(ctx); err != nil {
		return err
	}

	return nil
}
