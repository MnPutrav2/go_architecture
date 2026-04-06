package repository

import (
	"context"
	"database/sql"

	"github.com/MnPutrav2/go_architecture/internal/model"
	"github.com/MnPutrav2/go_architecture/pkg/query"
)

type UserRepository struct {
	db *sql.DB
}

func InituserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

// Entry
func (q *UserRepository) CreateUserRepository(ctx context.Context, request model.CreateUser) error {
	if err := query.Init[model.Users](q.db).Insert(request).Exec(ctx); err != nil {
		return err
	}

	return nil
}
