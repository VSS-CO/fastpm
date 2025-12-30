package downloader

import (
	"errors"
	"io"
	"net/http"
	"os"
)

func downloadURL(url, dest string) error {
    out, err := os.Create(dest)
    if err != nil {
        return err
    }
    defer out.Close()

    resp, err := http.Get(url)
    if err != nil {
        return err
    }
    defer resp.Body.Close()

    if resp.StatusCode != 200 {
        return errors.New("failed to download: " + resp.Status)
    }

    _, err = io.Copy(out, resp.Body)
    return err
}
