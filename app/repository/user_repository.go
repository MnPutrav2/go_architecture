package repository

import (
	"context"
	"database/sql"

	"github.com/MnPutrav2/go_architecture/app/model"
	"github.com/MnPutrav2/go_architecture/app/pkg/query"
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

func (q *UserRepository) GetUserRepository(ctx context.Context, username string) (model.Users, error) {
	result, err := query.Init[model.Users](q.db).Select("id, name, password, email, role").Where("name", username).Find(ctx)
	if err != nil {
		return model.Users{}, err
	}

	return result, nil
}

func (q *UserRepository) GetUserByIdRepository(ctx context.Context, id uuid.UUID) (model.Users, error) {
	result, err := query.Init[model.Users](q.db).Select("id, name, password, email, role").Where("id", id.String()).Find(ctx)
	if err != nil {
		return model.Users{}, err
	}

	return result, nil
}
