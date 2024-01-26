package hellsgopher

import (
	"path/filepath"
	"runtime"
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

const TH32CS_SNAPPROCESS = 0x00000002

// windows process type structure
type WinProcess struct {
	PID  int
	PPID int
	Exe  string // name of process
}

// return the pid(s) from the process name
func GetPidFromName(name string) ([]int, error) {
	var pids []int

	processes, err := ListAllProcesses()
	if err != nil {
		return pids, err
	}

	for _, proc := range processes {
		if proc.Exe == name {
			pids = append(pids, proc.PID)
		}
	}

	if len(pids) == 0 {
		return pids, ErrProcess_Not_Found
	} else {
		return pids, nil
	}
}

// get the name from the pid
func GetNameFromPid(pid int) (string, error) {
	processes, err := ListAllProcesses()
	if err != nil {
		return "", err
	}

	for _, proc := range processes {
		if proc.PID == pid {
			return proc.Exe, nil
		}
	}

	return "", ErrPid_Not_Found
}

// list all running processes
func ListAllProcesses() ([]WinProcess, error) {
	hSnapshot, err := windows.CreateToolhelp32Snapshot(TH32CS_SNAPPROCESS, 0)
	if err != nil {
		return nil, err
	}
	defer windows.CloseHandle(hSnapshot)

	var pEntry windows.ProcessEntry32
	pEntry.Size = uint32(unsafe.Sizeof(pEntry))

	err = windows.Process32First(hSnapshot, &pEntry)
	if err != nil {
		return nil, err
	}

	results := make([]WinProcess, 0, 50)
	for {
		results = append(results, newWinProcess(&pEntry))

		err = windows.Process32Next(hSnapshot, &pEntry)
		if err != nil {
			if err == windows.ERROR_NO_MORE_FILES {
				return results, nil
			}
			return nil, err
		}
	}
}

// get pid of current process
func GetCurrentPid() int {
	return syscall.Getpid()
}

// get ppid of current process
func GetCurrentPpid() int {
	return syscall.Getppid()
}

// get the path of the current process
func GetCurrentProcPath() (string, error) {
	n := uint32(1024)
	var buf []uint16
	for {
		buf = make([]uint16, n)
		r, err := windows.GetModuleFileName(0, &buf[0], n)
		if err != nil {
			return "", err
		}
		if r < n {
			break
		}
		n += 1024
	}
	return syscall.UTF16ToString(buf), nil
}

// get the name of the current process
func GetCurrentProcName() (string, error) {
	procPath, err := GetCurrentProcPath()
	if err != nil {
		return "", err
	}

	return filepath.Base(procPath), nil
}

// get the arch of the current process
func GetCurrentProcArch() string {
	return runtime.GOARCH
}

/*
AUXILIARY FUNCTION
*/

func newWinProcess(pe *windows.ProcessEntry32) WinProcess {
	end := 0
	for {
		if pe.ExeFile[end] == 0 {
			break
		}
		end++
	}

	return WinProcess{
		PID:  int(pe.ProcessID),
		PPID: int(pe.ParentProcessID),
		Exe:  windows.UTF16ToString(pe.ExeFile[:end]),
	}
}
