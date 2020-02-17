package utils

import "golang.org/x/crypto/bcrypt"

// HashPassword encrypts the password
func HashPassword(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return string(hash)
}

// VerifyPassword decrypts the password and verifies with input password
func VerifyPassword(hashPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
	if err != nil {
		return false
	}
	return true
}
