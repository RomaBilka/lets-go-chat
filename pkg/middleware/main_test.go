package middleware

import (
	"os"
	"testing"
)

var testLog *Log

func TestMain(m *testing.M) {
	testLog = newTestLogStdout()
	os.Exit(m.Run())
}

type Log struct {
	name     string
	messages map[string]string
}

func newTestLogStdout() *Log {
	return &Log{}
}

func (l *Log) Init(name string) {
	l.name = name
	l.messages = make(map[string]string)
}
func (l *Log) GetName() string{
	return l.name
}

func (l *Log) GetMessage(key string) string {
	return l.messages[key]
}

func (l *Log) AddMessage(key, value string) {
	l.messages[key] = value
}
func (l *Log) Print() {
}