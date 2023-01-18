package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var jwtKey = []byte("supersecretkey")

type JWTClaim struct {
	Username string `json:"user_username"`
	Email    string `json:"user_email"`
	jwt.RegisteredClaims
}

func GenerateJWT(email string, username string) (tokenString string, err error) {
	claims := &JWTClaim{
		Email:    email,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "Test",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(jwtKey)
	return
}

func ValidateToken(signedToken string) (err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		},
	)

	if err != nil {
		return
	}

	claims, ok := token.Claims.(*JWTClaim)
	if !ok {
		err = errors.New("couldn't parse claims")
		return
	}

	if claims.ExpiresAt.Unix() < time.Now().Local().Unix() {
		err = errors.New("token expired")
		return
	}

	return

}
