package hellsgopher

import (
	"errors"
	"net"
	"os"
	"runtime"
	"strings"

	"github.com/pbnjay/memory"
)

// check if known VM files exist on system
func VmCheckFiles() (bool, error) {
	var files []string = []string{"C:\\windows\\System32\\Drivers\\Vmmouse.sys", "C:\\windows\\System32\\Drivers\\vm3dgl.dll", "C:\\windows\\System32\\Drivers\\vmdum.dll", "C:\\windows\\System32\\Drivers\\vm3dver.dll", "C:\\windows\\System32\\Drivers\\vmtray.dll", "C:\\windows\\System32\\Drivers\\VMToolsHook.dll", "C:\\windows\\System32\\Drivers\\vmmousever.dll", "C:\\windows\\System32\\Drivers\\vmhgfs.dll", "C:\\windows\\System32\\Drivers\\vmGuestLib.dll", "C:\\windows\\System32\\Drivers\\VmGuestLibJava.dll", "C:\\windows\\System32\\Driversvmhgfs.dll", "C:\\windows\\System32\\Drivers\\VBoxMouse.sys", "C:\\windows\\System32\\Drivers\\VBoxGuest.sys", "C:\\windows\\System32\\Drivers\\VBoxSF.sys", "C:\\windows\\System32\\Drivers\\VBoxVideo.sys", "C:\\windows\\System32\\vboxdisp.dll", "C:\\windows\\System32\\vboxhook.dll", "C:\\windows\\System32\\vboxmrxnp.dll", "C:\\windows\\System32\\vboxogl.dll", "C:\\windows\\System32\\vboxoglarrayspu.dll", "C:\\windows\\System32\\vboxoglcrutil.dll", "C:\\windows\\System32\\vboxoglerrorspu.dll", "C:\\windows\\System32\\vboxoglfeedbackspu.dll", "C:\\windows\\System32\\vboxoglpackspu.dll", "C:\\windows\\System32\\vboxoglpassthroughspu.dll", "C:\\windows\\System32\\vboxservice.exe", "C:\\windows\\System32\\vboxtray.exe", "C:\\windows\\System32\\VBoxControl.exe"}

	isVm := false

	for _, path := range files {
		if _, err := os.Stat(path); err == nil {
			isVm = true
			return isVm, nil
		}
	}

	return false, nil
}

// check if known VM related processes are running
func VmCheckProcesses() (bool, error) {
	// get all processes
	processes, err := ListAllProcesses()
	if err != nil {
		return false, err
	}

	// check if any processes contain "vm" or "vbox"
	for _, proc := range processes {
		if strings.Contains(strings.ToLower(proc.Exe), "vm") || strings.Contains(strings.ToLower(proc.Exe), "vbox") {
			return true, nil
		}
	}

	return false, nil
}

// check whether the machine has less than or equal to x number of cores (default is 2, leave as 0 for default)
func VmCheckCores(count int) (bool, error) {
	if count < 0 {
		return false, errors.New("count must be positive")
	}

	cores := 2
	if count != 0 {
		cores = count
	}

	numOfProcessors := runtime.NumCPU()

	if cores <= numOfProcessors {
		return false, nil
	} else {
		return true, nil
	}
}

// check whether the machine has less than or equal to x mb of ram (default is 4196, leave as 0 for default)
func VmCheckRam(mb uint64) (bool, error) {
	var ram uint64 = 4196 * 1048576
	if mb != 0 {
		ram = mb * 1048576
	}

	amountOfRam := memory.TotalMemory()

	// check if ram is less than or equal to machine's ram
	if ram <= amountOfRam {
		return false, nil
	} else {
		return true, nil
	}
}

// check if machine can access 8.8.8.8
func VmCheckOnline() bool {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return false
	}
	defer conn.Close()

	return true
}
