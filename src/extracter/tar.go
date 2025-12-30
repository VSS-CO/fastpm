package extractor

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"os"
	"path/filepath"
)

func ExtractTarGz(src, dest string) error {
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

    tr := tar.NewReader(gz)
    for {
        hdr, err := tr.Next()
        if err == io.EOF {
            break
        }
        if err != nil {
            return err
        }

        target := filepath.Join(dest, hdr.Name)
        if hdr.Typeflag == tar.TypeDir {
            os.MkdirAll(target, 0755)
        } else {
            out, err := os.Create(target)
            if err != nil {
                return err
            }
            io.Copy(out, tr)
            out.Close()
        }
    }
    return nil
}
