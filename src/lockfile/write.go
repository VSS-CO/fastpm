package lockfile

import (
	"encoding/json"
	"os"
)

func WriteLockfile(path string, locks []Lock) error {
    f, err := os.Create(path)
    if err != nil {
        return err
    }
    defer f.Close()
    return json.NewEncoder(f).Encode(locks)
}
