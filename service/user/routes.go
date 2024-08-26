package user

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gitkoDev/medods_task/helpers"
	"github.com/gitkoDev/medods_task/models"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type Handler struct {
	storage models.UserStorage
}

func NewHandler(storage models.UserStorage) *Handler {
	return &Handler{
		storage: storage,
	}
}

func (h *Handler) RegisterRoutes(router *gin.Engine) {

	v1 := router.Group("/v1")
	{
		v1.GET("/health", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{"message": "okay"})
		})
		v1.GET("/auth", h.GetTokens)
		v1.POST("/auth/refresh", h.RefreshTokens)
	}

}

func (h *Handler) GetTokens(c *gin.Context) {
	user_ip := helpers.GetUserIp(c)
	user_id := c.Query("id")

	// Validate user id
	if user_id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "no guid provided"})
		return
	}

	tokens := helpers.GenerateTokens(user_id)

	if !h.storage.IsExistingUser(c, user_id) {
		h.storage.AddUser(c, user_id, user_ip, tokens.RefreshTokenHash)
	} else {
		h.storage.UpdateRefreshTokenHash(c, user_id, tokens.RefreshTokenHash)
	}

	c.JSON(http.StatusOK, gin.H{"access_token": tokens.AcessToken, "refresh_token": tokens.RefreshToken, "user_ip": user_ip})
}

func (h *Handler) RefreshTokens(c *gin.Context) {
	var tokenPair models.TokenPair
	user_ip := helpers.GetUserIp(c)

	if !h.storage.IsExistingIp(c, user_ip) {
		helpers.SendEmail(c)
		return
	}

	// Validate tokens
	if err := c.BindJSON(&tokenPair); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid JSON or token structure"})
		return
	}

	tokenClaims := &models.TokenClaims{}
	_, err := jwt.ParseWithClaims(tokenPair.AcessToken, tokenClaims, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid access token"})
		return
	} else if bcrypt.CompareHashAndPassword(tokenClaims.RefreshTokenHash, tokenPair.RefreshToken) != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid refresh token"})
		return
	} else if !h.storage.IsExistingHash(c, tokenClaims.RefreshTokenHash) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "no such refresh token in db"})
		return
	}

	tokens := helpers.GenerateTokens(tokenClaims.ID)

	c.JSON(http.StatusOK, gin.H{"access_token": tokens.AcessToken, "refresh_token": tokens.RefreshToken})
}
