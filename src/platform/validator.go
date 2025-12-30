package platform

func IsWindows() bool {
	return OS() == "windows"
}
