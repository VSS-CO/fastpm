package cache

import (
	"io"
	"os"
	"path/filepath"
)

func SaveFile(src, pkg string) error {
    dest := filepath.Join("./cache", pkg)
    if err := os.MkdirAll(filepath.Dir(dest), 0755); err != nil {
        return err
    }
    in, err := os.Open(src)
    if err != nil {
        return err
    }
    defer in.Close()
    out, err := os.Create(dest)
    if err != nil {
        return err
    }
    defer out.Close()
    _, err = io.Copy(out, in)
    return err
}
