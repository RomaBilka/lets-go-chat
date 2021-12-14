package hasher

import (
	"testing"

	"github.com/bxcodec/faker/v3"
	"golang.org/x/crypto/bcrypt"
)

var password string = "test_password"

func TestHashPassword(t *testing.T) {
	hash, _ := HashPassword(password)

	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)); err != nil {
		t.Errorf("Hash Password is wrong")
	}
}

func TestCheckPasswordHash(t *testing.T) {
	hash, _ := HashPassword(password)

	if ok := CheckPasswordHash(password, hash); !ok {
		t.Errorf("CheckPasswordHash works incorrectly")
	}
}

func TestCheckPasswordHashEmptyPassword(t *testing.T) {
	hash, _ := HashPassword(password)

	if ok := CheckPasswordHash("", hash); ok {
		t.Errorf("CheckPasswordHash works incorrectly")
	}
}

func BenchmarkCalculate(b *testing.B) {
	type Password struct {
		password string `faker:"password"`
	}
	p := Password{}

	if err := faker.FakeData(&p); err != nil {
		b.Errorf(err.Error())
	}

	for i := 0; i < b.N; i++ {
		HashPassword(p.password)
	}
}
