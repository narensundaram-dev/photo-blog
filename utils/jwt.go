package utils

import (
	"encoding/json"
	"strconv"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// JWTSecret to generate encrypted JWT token
var JWTSecret = []byte("yOuCaNThAckME")

// Jwt struct to manage the auth token
type Jwt struct {
	UID      uint
	Name     string
	Username string
}

// GenerateToken returns JWT Token
func (j *Jwt) GenerateToken() string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":      json.Number(strconv.FormatInt(time.Now().AddDate(0, 0, 1).Unix(), 10)),
		"iat":      json.Number(strconv.FormatInt(time.Now().Unix(), 10)),
		"uid":      j.UID,
		"name":     j.Name,
		"username": j.Username,
	})

	tokenStr, err := token.SignedString(JWTSecret)
	if err != nil {
		panic(err)
	}
	return tokenStr
}

// ValidateToken validates JWT Token
func (j *Jwt) ValidateToken(tokenStr string) bool {
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return JWTSecret, nil
	})
	if err != nil {
		return false
	} else if !token.Valid {
		return false
	} else {
		return true
	}
}
