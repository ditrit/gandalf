package models

import jwt "github.com/dgrijalva/jwt-go"

//Token struct declaration
type Claims struct {
	UserID uint
	Name   string
	Email  string
	Tenant string
	*jwt.StandardClaims
}
