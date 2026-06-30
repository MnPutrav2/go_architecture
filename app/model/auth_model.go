package model

import (
	"time"

	"github.com/google/uuid"
)

// Database model
type RefreshToken struct {
	ID        uuid.UUID `json:"id" db:"id" structure:"UUID;primary key;default;gen_random_uuid()"`
	UserID    uuid.UUID `json:"user_id" db:"user_id" structure:"UUID;not null" relation:"users.id;cascade"`
	TokenHash string    `json:"token_hash" db:"token_hash" structure:"varchar(255);not null"`
	ExpiredAt time.Time `json:"expired_at" db:"expired_at" structure:"timestamp;not null"`
}

// Database model

type Token struct {
	Token        string    `json:"token"`
	RefreshToken string    `json:"refresh_token"`
	Expired      time.Time `json:"expired"`
}

type RefreshTokenPayload struct {
	UserID    uuid.UUID `json:"user_id" db:"user_id"`
	TokenHash string    `json:"token_hash" db:"token_hash"`
	ExpiredAt time.Time `json:"expired_at" db:"expired_at"`
}

// Request

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" db:"token_hash" validate:"required"`
}
