package extractor

import (
	"compress/gzip"
	"io"
	"os"
)

func ExtractGzip(src, dest string) error {
    f, err := os.Open(src)
    if err != nil {
        return err
    }
    defer f.Close()

    gz, err := gzip.NewReader(f)
    if err != nil {
        return err
    }
    defer gz.Close()

    out, err := os.Create(dest)
    if err != nil {
        return err
    }
    defer out.Close()

    _, err = io.Copy(out, gz)
    return err
}
