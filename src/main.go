package main

import (
	"flag"
	"fmt"
	"os"

	"fastpm/src/build"
	"fastpm/src/logger"
	"fastpm/src/resolver"
)

func main() {
    cmd := flag.String("cmd", "install", "Command: install/remove/update/list/info")
    pkg := flag.String("pkg", "", "Package name")
    ver := flag.String("ver", "latest", "Package version")
    target := flag.String("target", "./node_modules", "Install directory")
    concurrency := flag.Int("c", 4, "Concurrent downloads")
    flag.Parse()

    switch *cmd {
    case "install":
        if *pkg == "" {
            fmt.Println("Package name required")
            os.Exit(1)
        }
        deps, err := resolver.ResolveDependencies(*pkg, *ver)
        if err != nil {
            logger.Error(fmt.Sprintf("Failed to resolve dependencies: %v", err))
            os.Exit(1)
        }
        err = build.InstallPackages(deps, *target, *concurrency)
        if err != nil {
            logger.Error(fmt.Sprintf("Installation failed: %v", err))
            os.Exit(1)
        }
        logger.Success("Installation complete!")
    case "remove":
        if *pkg == "" {
            fmt.Println("Package name required")
            os.Exit(1)
        }
        err := build.RemovePackage(*pkg, *target)
        if err != nil {
            logger.Error(fmt.Sprintf("Failed to remove: %v", err))
            os.Exit(1)
        }
        logger.Success("Removed package successfully")
    default:
        fmt.Println("Unknown command")
    }
}
