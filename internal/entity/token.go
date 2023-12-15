package entity

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type Token struct {
	UserId       int64     `json:"user_id"`
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
	ExpiresAt    time.Time `json:"expires_at"`
}

type TokenClaims struct {
	UserId int64 `json:"uid,omitempty"`
	jwt.RegisteredClaims
}
