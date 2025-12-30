package lockfile

import (
	"encoding/json"
	"os"
)

const DefaultLockfileName = "blaze.lock"

func ReadLockfile(path string) ([]Lock, error) {
	if path == "" {
		path = DefaultLockfileName
	}

	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var locks []Lock
	err = json.NewDecoder(f).Decode(&locks)
	return locks, err
}
