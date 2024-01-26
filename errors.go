package hellsgopher

import "errors"

var (
	ErrFunction_Not_Supported = errors.New("function currently not supported")       // function is not currently implemented
	ErrPid_Not_Found          = errors.New("pid not found")                          // pid was not found on system
	ErrProcess_Not_Found      = errors.New("a process with that name was not found") // process was not found on system
)
