package logger

import (
	"fmt"
	"os"
)

var file *os.File

const (
	ColorReset  = "\033[0m"
	ColorRed    = "\033[31m"
	ColorGreen  = "\033[32m"
	ColorYellow = "\033[33m"
	ColorCyan   = "\033[36m"
)

func InitLogger(filename string) {
	if filename != "" {
		var err error
		file, err = os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Println("err with creating file:", err)
		}
	}
}

func Success(msg string) {
	print(ColorGreen, msg)
}

func Error(msg string) {
	print(ColorRed, msg)
}

func Info(msg string) {
	print(ColorYellow, msg)
}

func Debug(msg string) {
	print(ColorCyan, msg)
}

func print(color, msg string) {
	fmt.Println(color + msg + ColorReset)
	if file != nil {
		file.WriteString(msg + "\n")
	}
}

func CloseLogger() {
	if file != nil {
		file.Close()
	}
}
