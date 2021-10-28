//Home task (retraining program)
package main

import (
	"fmt"
	"github.com/RomaBiliak/lets-go-chat/pkg/hasher"
)

func main() {
	password, confirm := "test", "test"
	hashPassword, err := hasher.HashPassword(password)
	if err == nil {
		fmt.Println(hasher.CheckPasswordHash(confirm, hashPassword))
	}
	fmt.Println(err)
}
