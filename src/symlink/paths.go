package symlink

import "path/filepath"

func NodeModulesPath(pkg string) string {
    return filepath.Join("./node_modules", pkg)
}
