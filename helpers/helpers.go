package helpers

import (
	"net/http"
	"net/smtp"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetUserIp(c *gin.Context) string {
	ip := c.Request.Header.Get("X-Real-Ip")
	if ip == "" {
		ip = c.Request.Header.Get("X-Forwarded-For")
	}
	if ip == "" {
		ip = c.Request.RemoteAddr
	}

	index := strings.Index(ip, ":")
	return ip[:index]

}

func SendEmail(c *gin.Context) {
	from := "mock.email@gmail.com"
	password := "secret_password1234"

	to := []string{
		"mock.recepient@gmail.com",
	}

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	auth := smtp.PlainAuth("", from, password, smtpHost)

	message := "Your data was accessed from a different ip. If it was not you, please contact our techical support team"

	smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, []byte(message))

	c.JSON(http.StatusUnauthorized, gin.H{"error": "Your ip is different, please visit your email and confirm your identity"})
}
