package extractor

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
)

func ExtractZip(src, dest string) error {
    r, err := zip.OpenReader(src)
    if err != nil {
        return err
    }
    defer r.Close()

    for _, f := range r.File {
        path := filepath.Join(dest, f.Name)
        if f.FileInfo().IsDir() {
            os.MkdirAll(path, 0755)
            continue
        }
        os.MkdirAll(filepath.Dir(path), 0755)
        out, err := os.Create(path)
        if err != nil {
            return err
        }
        rc, err := f.Open()
        if err != nil {
            out.Close()
            return err
        }
        _, err = io.Copy(out, rc)
        out.Close()
        rc.Close()
        if err != nil {
            return err
        }
    }
    return nil
}
