package models

import "github.com/dgrijalva/jwt-go"

type Claim struct {
	Username []string `json:"roles"`
	// recommended having
	jwt.StandardClaims
}