package model

import (
	"time"

	"github.com/google/uuid"
)

// Database model

type Users struct {
	ID       uuid.UUID `json:"id" db:"id" structure:"UUID;primary key;default;gen_random_uuid()"`
	Name     string    `json:"name" db:"name" structure:"varchar(255);not null"`
	Password string    `json:"password" db:"password" structure:"varchar(2000);not null"`
	Email    string    `json:"email" db:"email" structure:"varchar(100)not null"`
}

// Database model

type CreateUser struct {
	Name     string `json:"name" db:"name" validate:"required;min:3;max:5"`
	Password string `json:"password" db:"password" validate:"required;min:8"`
	Email    string `json:"email" db:"email" validate:"required"`
}

type LoginUser struct {
	Name     string `json:"name" db:"name" validate:"required;min:3;max:5"`
	Password string `json:"password" db:"password" validate:"required;min:8"`
}

type Token struct {
	Token   string    `json:"token"`
	Expired time.Time `json:"expired"`
}
