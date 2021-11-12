package models

type User struct {
	Id       uint64
	Name     string
	Password string
}

type UserRepository interface {
	GetUserByName(name string) (User, error)
	CreateUser(user User) (User, error)
}
