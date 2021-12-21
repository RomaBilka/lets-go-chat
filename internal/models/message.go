package models

type MessageId int

type Message struct {
	Id        MessageId
	UserId    UserId
	Message   string
	CreatedAt int
}
