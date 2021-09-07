package models

import jwt "github.com/golang-jwt/jwt/v4"

//Token struct declaration
type Claims struct {
	UserID uint
	Name   string
	Email  string
	Tenant string
	*jwt.StandardClaims
}
