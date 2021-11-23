package hasher

import (
	"testing"

	"golang.org/x/crypto/bcrypt"
)

var password string = "test_password"

func TestHashPassword(t *testing.T) {
	hash, _ := HashPassword(password)

	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	if err != nil {
		t.Errorf("Hash Password is wrong")
	}
}

func TestCheckPasswordHash(t *testing.T) {
	hash, _ := HashPassword(password)

	ok := CheckPasswordHash(password, hash)

	if !ok {
		t.Errorf("CheckPasswordHash works incorrectly")
	}
}

func TestCheckPasswordHashEmptyPassword(t *testing.T) {
	hash, _ := HashPassword(password)

	ok := CheckPasswordHash("", hash)

	if ok {
		t.Errorf("CheckPasswordHash works incorrectly")
	}
}
