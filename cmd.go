package hellsgopher

import (
	"fmt"
	"os/exec"
	"runtime"
)

// will run a command with either `bash -c` or `cmd /C` and return the output
func CmdReturn(command string) (string, error) {
	var cmd *exec.Cmd
	if runtime.GOOS == "linux" {
		cmd = exec.Command("bash", "-c", command)
	} else if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/C", command)
	} else {
		return "", ErrUnknownGOOS
	}

	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}

	return string(output), nil
}

// will run a command with either `bash -c` or `cmd /C` and print output and error to STDOUT
func CmdSTDOUT(command string) {
	var cmd *exec.Cmd
	if runtime.GOOS == "linux" {
		cmd = exec.Command("bash", "-c", command)
	} else if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/C", command)
	} else {
		Error("unknown GOOS")
		return
	}

	output, err := cmd.CombinedOutput()
	if err != nil {
		Error(err.Error())
		return
	}

	print(string(output))
}

// will run a command with either `bash -c` or `cmd /C` and will provide no output
func CmdNoOut(command string) {
	var cmd *exec.Cmd
	if runtime.GOOS == "linux" {
		cmd = exec.Command("bash", "-c", command)
	} else if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/C", command)
	} else {
		return
	}

	cmd.CombinedOutput()
}

// runs a command with `powershell -Command "& {}"` and returns the output as a string
func PSReturn(command string) (string, error) {
	if runtime.GOOS != "windows" {
		return "", ErrNotWin
	}

	cmd := exec.Command("powershell", "-Command", fmt.Sprintf("\"& {%s}\"", command))
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}

	return string(output), nil
}

// runs a command with `powershell -Command "& {}"` and prints the output and error to STDOUT
func PSSTDOUT(command string) {
	if runtime.GOOS != "windows" {
		return
	}

	cmd := exec.Command("powershell", "-Command", fmt.Sprintf("\"& {%s}\"", command))
	output, err := cmd.CombinedOutput()
	if err != nil {
		Error(err.Error())
	}

	print(string(output))
}

// runs a command with `powershell -Command "& {}"` and will provide no output
func PSNoOut(command string) {
	if runtime.GOOS != "windows" {
		return
	}

	cmd := exec.Command("powershell", "-Command", fmt.Sprintf("\"& {%s}\"", command))
	cmd.CombinedOutput()
}
