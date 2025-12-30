package cache

import "path/filepath"

func CachePath(pkg string) string {
    return filepath.Join("./cache", pkg)
}
