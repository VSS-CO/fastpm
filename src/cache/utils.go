package cache

import "os"

func EnsureCacheDir() error {
    return os.MkdirAll("./cache", 0755)
}
