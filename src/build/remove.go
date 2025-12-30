package build

import (
	"fastpm/src/logger"
	"fastpm/src/symlink"
	"os"
)

func RemovePackage(pkg, target string) error {
    path := target + "/" + pkg
    os.RemoveAll(path)
    symlink.Remove("./node_modules/" + pkg)
    logger.Info("Removed " + pkg)
    return nil
}
