package myJwt

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/peetwerapat/learnhub-go-api/internal/domain"
)

var jwtKey = []byte(GetJWTSecret())

func GetJWTSecret() string {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "default_secret"
	}
	return secret
}

type Claims struct {
	ID    uint   `json:"id"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func CreateToken(user *domain.User, expiration time.Duration) (string, error) {
	claims := &Claims{
		ID:    user.ID,
		Email: user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiration)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(GetJWTSecret()))
}
