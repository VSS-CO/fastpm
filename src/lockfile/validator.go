package lockfile

func Validate(locks []Lock) bool {
	for _, l := range locks {
		if l.Name == "" || l.Version == "" || l.URL == "" {
			return false
		}
	}
	return true
}
