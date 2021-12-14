package log

import (
	"fmt"
	"io"
	"log"
	"os"
)

type Log struct {
	name     string
	messages map[string]string
}

func NewLogStdout() *Log {
	return &Log{}
}

func (l *Log) Init(name string) {
	l.name = name
	l.messages = make(map[string]string)
}

func (l *Log) AddMessage(key, value string) {
	l.messages[key] = value
}

func (l *Log) Print() {
	_, err := io.WriteString(os.Stdout, fmt.Sprintf("========= Start: %s =========\n", l.name))

	for key, value := range l.messages {
		_, err = io.WriteString(os.Stdout, fmt.Sprintf("%s: %s\n", key, value))
	}

	_, err = io.WriteString(os.Stdout, fmt.Sprintf("========= End: %s =========\n", l.name))

	if err != nil {
		log.Fatal(err)
	}
}
