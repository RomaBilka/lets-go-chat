package log

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

type Log struct {
	name     string
	messages map[string]string
}

func NewLogStdout(name string) *Log {
	return &Log{name: name, messages: make(map[string]string)}
}

func (l *Log) AddMessage(key, value string) {
	l.messages[key] = value
}

func (l *Log) Print() {
	w := bufio.NewWriter(os.Stdout)
	fmt.Fprint(w, "Hello, ")
	_, err := io.WriteString(os.Stdout, fmt.Sprintf("========= Start: %s =========\n", l.name))

	for key, value := range l.messages {
		_, err = io.WriteString(os.Stdout, fmt.Sprintf("%s: %s\n", key, value))
	}

	_, err = io.WriteString(os.Stdout, fmt.Sprintf("========= End: %s =========\n", l.name))
	w.Flush()
	if err != nil {
		log.Fatal(err)
	}
}
