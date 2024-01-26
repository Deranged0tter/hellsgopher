package hellsgopher

import (
	"os/exec"
	"syscall"

	"golang.org/x/sys/windows"
)

// will run cmd.exe and return output
func CmdReturn(command string) (string, error) {
	cmdPath := "C:\\Windows\\system32\\cmd.exe"
	cmdInstance := exec.Command(cmdPath, "/c", command)
	cmdInstance.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	cmdOut, err := cmdInstance.Output()

	return string(cmdOut), err
}

// will run cmd.exe and print output to STDOUT
func CmdStdOUT(command string) {
	cmdPath := "C:\\Windows\\system32\\cmd.exe"
	cmdInstance := exec.Command(cmdPath, "/c", command)
	cmdInstance.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	cmdOut, err := cmdInstance.Output()
	if err != nil {
		print(err)
		return
	}

	print(string(cmdOut))
}

// will run cmd.exe and provide no output
func CmdNoOut(command string) {
	cmdPath := "C:\\Windows\\system32\\cmd.exe"
	cmdInstance := exec.Command(cmdPath, "/c", command)
	cmdInstance.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	cmdInstance.Output()
}

// will run powershell command and return output
func PsReturn(command string) (string, error) {
	psPath := "C:\\Windows\\System32\\WindowsPowerShell\\v1.0\\powershell.exe"
	psInstance := exec.Command(psPath, command)
	psInstance.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	psOut, err := psInstance.Output()

	return string(psOut), err
}

// will run powershell command and print output to STDOUT
func PsStdOut(command string) {
	psPath := "C:\\Windows\\System32\\WindowsPowerShell\\v1.0\\powershell.exe"
	psInstance := exec.Command(psPath, command)
	psInstance.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	psOut, err := psInstance.Output()
	if err != nil {
		print(err)
		return
	}

	print(string(psOut))
}

// will run powershell command and provide no output
func PsNoOut(command string) {
	psPath := "C:\\Windows\\System32\\WindowsPowerShell\\v1.0\\powershell.exe"
	psInstance := exec.Command(psPath, command)
	psInstance.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	psInstance.Output()
}

// will run powershell command and return output (with token)
func PsReturnT(command string, token windows.Token) (string, error) {
	psPath := "C:\\Windows\\System32\\WindowsPowerShell\\v1.0\\powershell.exe"
	psInstance := exec.Command(psPath, command)
	psInstance.SysProcAttr = &syscall.SysProcAttr{HideWindow: true, Token: syscall.Token(token)}
	psOut, err := psInstance.Output()

	return string(psOut), err
}

// will run powershell command and print output to STDOUT (with token)
func PsStdOutT(command string, token windows.Token) {
	psPath := "C:\\Windows\\System32\\WindowsPowerShell\\v1.0\\powershell.exe"
	psInstance := exec.Command(psPath, command)
	psInstance.SysProcAttr = &syscall.SysProcAttr{HideWindow: true, Token: syscall.Token(token)}
	psOut, err := psInstance.Output()
	if err != nil {
		print(err)
		return
	}

	print(string(psOut))
}

// will run powershell command and provide no output (with token)
func PsNoOutT(command string, token windows.Token) {
	psPath := "C:\\Windows\\System32\\WindowsPowerShell\\v1.0\\powershell.exe"
	psInstance := exec.Command(psPath, command)
	psInstance.SysProcAttr = &syscall.SysProcAttr{HideWindow: true, Token: syscall.Token(token)}
	psInstance.Output()
}
