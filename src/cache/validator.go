package cache

import "os"

func Exists(pkg string) bool {
    _, err := os.Stat("./cache/" + pkg)
    return err == nil
}
