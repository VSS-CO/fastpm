package services

import "fmt"

func RecordDownload(pkg string, size int64) {
    fmt.Printf("Downloaded %s: %d bytes\n", pkg, size)
}
