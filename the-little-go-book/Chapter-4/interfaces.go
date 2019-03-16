package main

import (
	"fmt"
)

type Logger interface {
	Log(message string)
}

type ConsoleLogger struct {}
func (l ConsoleLogger) Log(message string) {
	fmt.Println(message)
}

// ConsoleLogger is a Logger because it implements the interface
func getNewLogger() Logger {
	return &ConsoleLogger{}
}

func main() {
	logger := getNewLogger()
	logger.Log("hey")
}
