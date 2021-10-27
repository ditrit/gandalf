package models

import (
	jwt "github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

//Token struct declaration
type Claims struct {
	UserID uuid.UUID
	Name   string
	Email  string
	Tenant string
	*jwt.StandardClaims
}
