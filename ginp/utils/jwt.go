package utils

import (
	"ginp/models"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("a_secret_create")

// Claims 声明
type Claims struct {
	UserID uint
	jwt.StandardClaims
}

// ReleaseToken 发放token
func ReleaseToken(user models.User) (string, error) {
	// expirtionTime := time.Now().Add(7 * 24 * time.Hour)
	expirtionTime := time.Now().Add(30 * time.Second)
	claims := &Claims{
		UserID: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirtionTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "one user",
			Subject:   "user token",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// ParseToken 解析token
func ParseToken(tokenstring string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenstring, claims, func(token *jwt.Token) (i interface{}, err error) {
		return jwtKey, nil
	})

	return token, claims, err
}
