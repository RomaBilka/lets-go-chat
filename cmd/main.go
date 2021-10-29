//Home task (retraining program)
package main

import (
	"fmt"

	"github.com/RomaBiliak/lets-go-chat/pkg/hasher"
)

func main() {
	testComparePassword()
}

//testComparePassword Password comparison test function
func testComparePassword() {
	password := "test"

	hashPassword, err := hasher.HashPassword(password)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	ok := hasher.CheckPasswordHash(password, hashPassword)
	if !ok {
		fmt.Println("error verifying password")
		return
	}

	fmt.Println("password matches")
}
