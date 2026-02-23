package userRepository

import (
	"context"
	"database/sql"

	userModel "github.com/MnPutrav2/go_architecture/internal/model/user"
)

type userRepository struct {
	db *sql.DB
}

type UserRepository interface {
	CreateUser(ctx context.Context, user userModel.Create) error
}

func InitUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

// Entry

func (q *userRepository) CreateUser(ctx context.Context, user userModel.Create) error {

	if _, err := q.db.ExecContext(ctx, create, user.Name); err != nil {
		return err
	}

	return nil
}
