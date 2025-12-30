package build

func UpdatePackage(pkg, target string) error {
	RemovePackage(pkg, target)
	return nil
}
