package resolver

func Validate(pkg Node) bool {
	return pkg.Name != "" && pkg.Version != ""
}
