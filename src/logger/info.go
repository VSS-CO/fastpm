package logger

import "fmt"

func Info(msg string) {
    fmt.Println("\033[34m[INFO]\033[0m", msg)
}
