package logger

import "fmt"

func Success(msg string) {
    fmt.Println("\033[32m[SUCCESS]\033[0m", msg)
}
