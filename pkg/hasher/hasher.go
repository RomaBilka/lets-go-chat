package hasher

import (
	"golang.org/x/crypto/bcrypt"
)

// HashPassword returns the bcrypt hash of the password and error
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 5)
	return string(bytes), err
}

// CheckPasswordHash compares a password and bcrypt hashed, returns a Boolean value
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
