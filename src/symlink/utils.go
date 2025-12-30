package symlink

import "os"

func EnsureDir(dir string) error {
    return os.MkdirAll(dir, 0755)
}
