package utils

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/lailiseptiandi/api-user-auth/config"
)

func GetSecret() []byte {
	config.LoadEnv()
	var JWT_SECRET_KEY = []byte(os.Getenv("JWT_SECRET_KEY"))
	return JWT_SECRET_KEY
}

func GenerateToken(payload interface{}) (string, error) {

	tokenTime, _ := strconv.Atoi(os.Getenv("JWT_EXPIRED_TOKEN"))

	claim := jwt.MapClaims{}
	claim["sub"] = payload
	claim["exp"] = time.Now().Add(time.Minute * time.Duration(tokenTime)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	signedToken, err := token.SignedString(GetSecret())
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func ValidateToken(encodedToken string) (interface{}, error) {

	parsedToken, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("invalid token")
		}
		return GetSecret(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("validate: %w", err)
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok || !parsedToken.Valid {
		return nil, fmt.Errorf("validate: invalid token")
	}

	return claims["sub"], nil
}
