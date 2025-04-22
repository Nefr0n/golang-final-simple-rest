package dto

import (
	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	UserID int64 `json:"user_id"`
	jwt.RegisteredClaims
}
