package resolver

import (
	"fastpm/src/downloader"
	"fmt"
)

// FetchPackage fetches package metadata from the npm registry
func FetchPackage(name, version string) downloader.Package {
	if version == "" || version == "latest" {
		version = "latest"
	}

	// npm registry URL format
	url := fmt.Sprintf("https://registry.npmjs.org/%s/-/%s-%s.tgz", name, name, version)
	
	return downloader.Package{
		Name: name,
		URL:  url,
	}
}
