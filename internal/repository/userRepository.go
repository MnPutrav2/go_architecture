package repository

import (
	"database/sql"
)

type UserRepository struct {
	db *sql.DB
}

func InituserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

// Entry
	