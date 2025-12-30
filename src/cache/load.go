package cache

import "os"

func LoadFile(pkg string) (*os.File, error) {
    path := "./cache/" + pkg
    return os.Open(path)
}
