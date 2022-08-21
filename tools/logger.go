package tools

import "log"

type SomeLogger interface {
	Info(string)
	Error(string2 string)
}

type MyLogger struct{}

func (ml *MyLogger) Info(msg string) {
	log.Println("INFO:", msg)
}

func (ml *MyLogger) Error(msg string) {
	log.Println("ERROR:", msg)
}
