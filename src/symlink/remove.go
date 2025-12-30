package symlink

import "os"

func Remove(dest string) error {
    return os.Remove(dest)
}
