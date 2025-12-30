package lockfile

func Exists(locks []Lock, name string) bool {
	for _, l := range locks {
		if l.Name == name {
			return true
		}
	}
	return false
}
