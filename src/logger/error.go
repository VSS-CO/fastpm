package logger

import "fmt"

func Error(msg string) {
    fmt.Println("\033[31m[ERROR]\033[0m", msg)
}
