package models

type UserId int

type User struct {
	Id       UserId
	Name     string
	Password string
}
