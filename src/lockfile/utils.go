package lockfile

import "fmt"

func PrintLocks(locks []Lock) {
    for _, l := range locks {
        fmt.Printf("%s@%s -> %s\n", l.Name, l.Version, l.URL)
    }
}
