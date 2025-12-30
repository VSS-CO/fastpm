package logger

import "fmt"

func Warn(msg string) {
    fmt.Println("\033[33m[WARN]\033[0m", msg)
}
