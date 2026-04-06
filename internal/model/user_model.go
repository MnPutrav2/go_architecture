package model

import "github.com/google/uuid"

// Database model

type Users struct {
	ID   uuid.UUID `json:"id" db:"id" structure:"UUID;primary key;default;gen_random_uuid()"`
	Name string    `json:"name" db:"name" structure:"varchar(255);not null"`
}

// Database model

type CreateUser struct {
	Name string `json:"name" db:"name"`
}
