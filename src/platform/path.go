package platform

import "path/filepath"

func Join(parts ...string) string {
    return filepath.Join(parts...)
}
