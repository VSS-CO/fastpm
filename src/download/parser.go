package downloader

import "strings"

func ParseURL(url string) (string, string) {
    parts := strings.Split(url, "/")
    file := parts[len(parts)-1]
    return url, "./cache/" + file
}
