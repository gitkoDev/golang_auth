package helpers

import (
	"crypto/rand"
	"log"

	"github.com/gitkoDev/medods_task/models"
	"github.com/golang-jwt/jwt/v4"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

func GenerateTokens(uuid string) models.TokenPair {
	refreshToken, refreshTokenHash, err := GenerateRefreshToken()
	if err != nil {
		logrus.Fatal(err)
	}

	accessToken, err := GenerateAccessToken(refreshTokenHash, uuid)
	if err != nil {
		log.Fatal(err)
	}

	return models.TokenPair{refreshToken, refreshTokenHash, accessToken}
}

func GenerateAccessToken(hash []byte, uuid string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, &models.TokenClaims{
		Id: uuid, RefreshTokenHash: hash,
	})

	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}
	return tokenString, nil

}

func GenerateRefreshToken() ([]byte, []byte, error) {
	randomSequence := make([]byte, 32)
	if _, err := rand.Read(randomSequence); err != nil {
		return nil, nil, err
	}

	refreshTokenHash, err := bcrypt.GenerateFromPassword(randomSequence, bcrypt.DefaultCost)
	if err != nil {
		return nil, nil, err
	}

	return randomSequence, refreshTokenHash, nil
}
