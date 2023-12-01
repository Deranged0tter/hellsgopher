package hellsgopher

import (
	"os"
	"runtime"
)

// return the computer's hostname
func GetHostname() string {
	hostname, err := os.Hostname()
	if err != nil {
		return ""
	}

	return hostname
}

// get the domain name of the computer
func GetDomainName() string {
	return ""
}

// get the current operating system
func GetOS() string {
	return runtime.GOOS
}

// get the os build
func GetOSBuild() string {
	if GetOS() == "windows" {
		return ""
	} else if GetOS() == "linux" {
		return ""
	} else {
		return ""
	}
}

func GetOSVersion() string {
	return ""
}
