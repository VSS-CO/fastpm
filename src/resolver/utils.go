package resolver

import "fastpm/src/downloader"

func ResolveDependencies(name, version string) ([]downloader.Package, error) {
	pkg := FetchPackage(name, version)
	return []downloader.Package{pkg}, nil
}
