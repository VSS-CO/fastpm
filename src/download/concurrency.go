package downloader

import (
	"fmt"
	"sync"
	"time"
)

type Package struct {
    Name string
    URL  string
}

func DownloadAll(pkgs []Package, concurrency int) error {
    tasks := make(chan Package, len(pkgs))
    for _, p := range pkgs {
        tasks <- p
    }
    close(tasks)

    wg := sync.WaitGroup{}
    for i := 0; i < concurrency; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            for pkg := range tasks {
                err := Retry(3, 500*time.Millisecond, func() error {
                    return downloadURL(pkg.URL, "./cache/"+pkg.Name)
                })
                if err != nil {
                    fmt.Println("Download failed:", pkg.Name, err)
                } else {
                    fmt.Println("Downloaded:", pkg.Name)
                }
            }
        }()
    }
    wg.Wait()
    return nil
}
