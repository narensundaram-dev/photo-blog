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
	token := j.GetTokenMust(tokenStr)

	if !token.Valid {
		return false
	} else {
		return true
	}
}

// GetTokenMust returns Token struct or panics if any error
func (j *Jwt) GetTokenMust(tokenStr string) *jwt.Token {
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return JWTSecret, nil
	})
	if err != nil {
		panic(err)
	}
	return token
}

// GetUserInfo parses the user info from token and returns the same
func (j *Jwt) GetUserInfo(tokenStr string) *Jwt {
	token := j.GetTokenMust(tokenStr)
	claims := token.Claims.(jwt.MapClaims)

	if userID, ok := claims["uid"].(float64); ok {
		j.UID = uint(userID)
	} else {
		panic("invalid token: uid in token is not of a type uint convertible")
	}

	if userName, ok := claims["username"].(string); ok {
		j.Username = userName
	} else {
		panic("invalid token: username in token is not of a type string")
	}

	if name, ok := claims["name"].(string); ok {
		j.Name = name
	} else {
		panic("invalid token: name in token is not of a type string")
	}
	return j
}
