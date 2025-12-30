package platform

import "runtime"

func Arch() string {
	return runtime.GOARCH
}
