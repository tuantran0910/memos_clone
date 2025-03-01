package version

var Version = "0.24.1"

var DevVersion = "0.24.1"

func GetCurrentVersion(mode string) string {
	if mode == "dev" || mode == "demo" {
		return DevVersion
	}
	return Version
}
