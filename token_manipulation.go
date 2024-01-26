package hellsgopher

import (
	"github.com/fourcorelabs/wintoken"
	"golang.org/x/sys/windows"
)

// get the token from the current process
func GetCurrentToken() (windows.Token, error) {
	return GetTokenFromPid(0)
}

// get the token from a process given its pid
func GetTokenFromPid(pid int) (windows.Token, error) {
	token, err := wintoken.OpenProcessToken(pid, wintoken.TokenPrimary)
	if err != nil {
		return token.Token(), err
	}
	defer token.Close()

	return token.Token(), nil
}

// get the token from a process given its process name
func GetTokenFromName(procName string) (windows.Token, error) {
	var token windows.Token

	pids, err := GetPidFromName(procName)
	if err != nil {
		return token, err
	}

	return GetTokenFromPid(pids[0])
}
