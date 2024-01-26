package hellsgopher

import (
	"fmt"
	"runtime"
	"strings"
	"syscall"

	"golang.org/x/sys/windows"
)

const VER_NT_WORKSTATION = 0x0000001

// return the machine's hostname
func GetHostname() (string, error) {
	const format = windows.ComputerNamePhysicalDnsHostname

	var n uint32 = 64
	for {
		b := make([]uint16, n)

		err := windows.GetComputerNameEx(format, &b[0], &n)
		if err == nil {
			return syscall.UTF16ToString(b[:n]), nil
		}
		if err != syscall.ERROR_MORE_DATA {
			return "", err
		}

		if n <= uint32(len(b)) {
			return "", err
		}
	}
}

// return the domain name of the machine
func GetDomainName() (string, error) {
	return "", ErrFunction_Not_Supported
}

// return the machine's OS
func GetOS() string {
	return runtime.GOOS
}

// return the machine's OS Build Number
func GetOSBuild() string {
	osVersion := windows.RtlGetVersion()
	return fmt.Sprint(osVersion.BuildNumber)
}

// return the machine's OS Version
func GetOSVersion() string {
	osVersion := windows.RtlGetVersion()
	var osName string
	if osVersion.MajorVersion == 6 {
		switch osVersion.MinorVersion {
		case 0:
			if osVersion.ProductType == VER_NT_WORKSTATION {
				osName = "Vista"
			} else {
				osName = "Server 2008"
			}
		case 1:
			if osVersion.ProductType == VER_NT_WORKSTATION {
				osName = "7"
			} else {
				osName = "Server 2008 R2"
			}
		case 2:
			if osVersion.ProductType == VER_NT_WORKSTATION {
				osName = "8"
			} else {
				osName = "Server 2012"
			}
		case 3:
			if osVersion.ProductType == VER_NT_WORKSTATION {
				osName = "8.1"
			} else {
				osName = "Server 2012 R2"
			}
		}
	} else {
		if osVersion.ProductType == VER_NT_WORKSTATION {
			if strings.HasPrefix(fmt.Sprint(osVersion.BuildNumber), "22") {
				osName = "11"
			} else {
				osName = "10"
			}
		} else {
			osName = "Server 2016"
		}
	}
	return "Windows " + osName
}
