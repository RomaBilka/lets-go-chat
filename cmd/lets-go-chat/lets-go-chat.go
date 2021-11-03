//Home task (retraining program)
package main

import (
	"fmt"

	"github.com/RomaBiliak/lets-go-chat/internal/server"
	"github.com/RomaBiliak/lets-go-chat/pkg/hasher"
)

var ser = server.Server{}

func main() {
	testComparePassword()
	ser.Run(":3000")
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
		fmt.Println("error: verifying password")
		return
	}

	fmt.Println("password matches")
}
