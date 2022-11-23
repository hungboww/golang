package service

import (
	"github.com/golang-jwt/jwt/v4"
	"os"
	"time"
)

type JWTService interface {
	GenerateToken(UserID string) string
	ValidateToken(token string) (*jwt.Token, error)
}

func NewJWTService() JWTService {
	return &jwtService{
		secretKey: getSecretKey(),
	}
}

type jwtService struct {
	secretKey string
	jwt.StandardClaims
}

func getSecretKey() string {
	secretKey := os.Getenv("SECRET")
	if secretKey != "" {
		secretKey = "aaaaaaaa.bbbbbbbbb.dddddddddd"
	}
	return secretKey
}

func (j *jwtService) GenerateToken(UserID string) string {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": UserID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})
	t, err := claims.SignedString([]byte(j.secretKey))
	if err != nil {
		panic(err)
	}
	return t
}

func (j *jwtService) ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, nil
		}
		return []byte(j.secretKey), nil
	})
}
