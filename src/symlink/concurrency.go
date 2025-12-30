package symlink

import "sync"

func BatchCreate(pairs map[string]string) {
    var wg sync.WaitGroup
    for src, dest := range pairs {
        wg.Add(1)
        go func(s, d string) {
            defer wg.Done()
            Create(s, d)
        }(src, dest)
    }
    wg.Wait()
}
