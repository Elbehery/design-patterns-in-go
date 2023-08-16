package main

import (
	"fmt"
	"os"
)

type Logger interface {
	Log(msg string)
}

type ConsoleLogger struct{}

func (cl *ConsoleLogger) Log(msg string) {
	fmt.Println(msg)
}

type FileLogger struct {
	file *os.File
}

func (fl *FileLogger) Log(msg string) {
	fmt.Fprintln(fl.file, msg)
}

func CreateLogger(loggerType string) (Logger, error) {
	switch loggerType {
	case "console":
		logger := &ConsoleLogger{}
		return logger, nil
	case "file":
		f, err := os.Create("log.txt")
		if err != nil {
			return nil, err
		}
		return &FileLogger{f}, nil
	default:
		return nil, fmt.Errorf("unsupported logger")
	}
}

func main() {
	logger, err := CreateLogger("console")
	if err != nil {
		panic(err)
	}
	logger.Log("Hello from console logger!")

	logger, err = CreateLogger("file")
	if err != nil {
		panic(err)
	}
	logger.Log("Hello from file logger!")
}
