package models

import "time"

type MessageId int

type Message struct {
	Id        MessageId
	UserId    UserId
	Message   string
	CreatedAt time.Time
}
