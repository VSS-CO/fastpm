package cache

import (
	"os"
	"path/filepath"
)

func Cleanup(maxFiles int) error {
    files, err := filepath.Glob("./cache/*")
    if err != nil {
        return err
    }
    if len(files) <= maxFiles {
        return nil
    }
    for i := 0; i < len(files)-maxFiles; i++ {
        os.RemoveAll(files[i])
    }
    return nil
}
