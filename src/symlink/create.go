package symlink

import "os"

func Create(src, dest string) error {
    return os.Symlink(src, dest)
}
