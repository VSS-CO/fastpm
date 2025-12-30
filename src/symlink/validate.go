package symlink

import "os"

func Exists(dest string) bool {
    _, err := os.Lstat(dest)
    return err == nil
}
