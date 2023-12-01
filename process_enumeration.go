package hellsgopher

import (
	"os"
	"path/filepath"
	"runtime"

	"github.com/mitchellh/go-ps"
)

// return a map of all the current running processes
func ListProcesses() (map[int]string, error) {
	processMap := make(map[int]string)
	processList, err := ps.Processes()
	if err != nil {
		return nil, err
	}

	for i := range processList {
		process := processList[i]
		processMap[process.Pid()] = process.Executable()
	}

	return processMap, nil
}

// get the pid of the current process
func GetCurrentProcessID() int {
	return os.Getpid()
}

// get the ppid of the current process
func GetCurrentParentProcessID() int {
	return os.Getppid()
}

// return the proesses path of the current process
func GetCurrentProcessPath() (string, error) {
	return os.Executable()
}

// return the name of the current process
func GetCurrentProcessName() string {
	processPath, _ := GetCurrentProcessPath()

	return filepath.Base(processPath)
}

// return the arch of the running process
func GetCurrentProcessArch() string {
	return runtime.GOARCH
}
