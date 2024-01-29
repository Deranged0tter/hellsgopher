package hellsgopher

import (
	"fmt"
	"runtime"
	"strings"
	"syscall"
	"time"

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

// return the uptime of the machine in seconds
func Uptime() int {
	kernel32 := windows.NewLazySystemDLL("kernel32")
	GetTicketCount64 := kernel32.NewProc("GetTickCount64")

	r1, _, _ := GetTicketCount64.Call(0, 0, 0, 0)

	uptime := time.Duration(r1) * time.Millisecond

	return int(uptime.Seconds())
}

// return a slice of pipes on system
func GetPipes() ([]string, error) {
	var pipes []string
	pipePref := `\\.\pipe\`
	var data windows.Win32finddata

	handle, err := windows.FindFirstFile(windows.StringToUTF16Ptr(pipePref+"*"), &data)
	if err != nil {
		return nil, err
	}
	defer windows.FindClose(handle)

	for {
		pipeName := windows.UTF16ToString(data.FileName[:])
		pipes = append(pipes, pipeName)

		if err := windows.FindNextFile(handle, &data); err != nil {
			if err == windows.ERROR_NO_MORE_FILES {
				break
			}

			return pipes, err
		}
	}

	return pipes, nil
}
