package downloader

import "os"

func EnsureDir(dir string) error {
    if _, err := os.Stat(dir); os.IsNotExist(err) {
        return os.MkdirAll(dir, 0755)
    }
    return nil
}
