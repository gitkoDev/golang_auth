package models

import "github.com/golang-jwt/jwt/v4"

type TokenClaims struct {
	jwt.RegisteredClaims
	Id               string `json:"id"`
	RefreshTokenHash []byte `json:"refresh_token_hash"`
}

type TokenPair struct {
	RefreshToken     []byte `json:"refresh_token"`
	RefreshTokenHash []byte
	AcessToken       string `json:"access_token"`
}
