package hasher

import (
	"fmt"
	"testing"

	"github.com/bxcodec/faker/v3"
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

func BenchmarkCalculate(b *testing.B) {
	type Password struct {
		password string `faker:"password"`
	}
	p := Password{}
	err := faker.FakeData(&p)

	if err != nil {
		b.Errorf(err.Error())
	}
	fmt.Println(p.password)

	for i := 0; i < b.N; i++ {
		HashPassword(p.password)
	}
}
