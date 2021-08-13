package util

import (
	"log"
	"os"
)

func Info(message string) {
	Info := log.New(os.Stdout, "\u001b[34mINFO: \u001B[0m", log.LstdFlags|log.Lshortfile)
	Info.Println(message)
}

func Warning(message string) {
	Warning := log.New(os.Stdout, "\u001b[33mWARNING: \u001B[0m", log.LstdFlags|log.Lshortfile)
	Warning.Println(message)
}

func Debug(message string) {
	Debug := log.New(os.Stdout, "\u001b[36mDEBUG: \u001B[0m", log.LstdFlags|log.Lshortfile)
	Debug.Println(message)
}

func Error(message string, err error) {
	Error := log.New(os.Stdout, "\u001b[31mERROR: \u001b[0m", log.LstdFlags|log.Lshortfile)
	Error.Println(message, err)
}

func Fatal(message string, err error) {
	Error := log.New(os.Stdout, "\u001b[31mERROR: \u001b[0m", log.LstdFlags|log.Lshortfile)
	Error.Fatal(message, err)
}
