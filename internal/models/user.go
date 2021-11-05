package models

type User struct {
	Id       uint64
	Name     string
	Password string
}

var UserId uint64 = 0

var Users[]User