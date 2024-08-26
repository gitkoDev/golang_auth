package models

import (
	"github.com/gin-gonic/gin"
)

type User struct {
	Id               string `json:"id"`
	RefreshTokenHash string `json:"refresh_token_hash"`
}

type UserStorage interface {
	UpdateRefreshTokenHash(*gin.Context, string, []byte)
	IsExistingUser(*gin.Context, string) bool
	IsExistingHash(*gin.Context, []byte) bool
	IsExistingIp(*gin.Context, string) bool
	AddUser(*gin.Context, string, string, []byte)
}
