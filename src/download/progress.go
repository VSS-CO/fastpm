package downloader

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type ProgressWriter struct {
    total   int64
    current int64
}

func (pw *ProgressWriter) Write(p []byte) (int, error) {
    n := len(p)
    pw.current += int64(n)
    fmt.Printf("\rDownloading... %d/%d bytes", pw.current, pw.total)
    return n, nil
}

func DownloadWithProgress(url, dest string) error {
    resp, err := http.Get(url)
    if err != nil {
        return err
    }
    defer resp.Body.Close()

    out, err := os.Create(dest)
    if err != nil {
        return err
    }
    defer out.Close()

    pw := &ProgressWriter{total: resp.ContentLength}
    _, err = io.Copy(io.MultiWriter(out, pw), resp.Body)
    fmt.Println()
    return err
}
