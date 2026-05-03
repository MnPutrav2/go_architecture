package repository

import (
	"context"
	"database/sql"

	"github.com/MnPutrav2/go_architecture/internal/model"
	"github.com/MnPutrav2/go_architecture/pkg/query"
	"github.com/google/uuid"
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

func (q *UserRepository) GetUserRepository(ctx context.Context) ([]model.Users, error) {
	result, err := query.Init[model.Users](q.db).Select("id, name").FindAll(ctx)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (q *UserRepository) DeleteUserRepository(ctx context.Context, id uuid.UUID) error {
	if err := query.Init[model.Users](q.db).Delete("id", id.String()).Exec(ctx); err != nil {
		return err
	}

	return nil
}
