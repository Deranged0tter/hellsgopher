package hellsgopher

import "errors"

var (
	ErrFunction_Not_Supported = errors.New("function currently not supported")                   // function is not currently implemented
	ErrPid_Not_Found          = errors.New("pid not found")                                      // pid was not found on system
	ErrProcess_Not_Found      = errors.New("a process with that name was not found")             // process was not found on system
	ErrFile_Not_Found         = errors.New("the provided file path was not found on the system") // a provided path was not found
	ErrFile_Not_DLL           = errors.New("the provided file is not a dll")                     // file needed is a dll and provided file is not a dll
)
