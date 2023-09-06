# hellsgopher
Go library for malware development

# THIS IS CURRENTLY UNDER DEVELOPMENT

# Functions
## Command Line Functions
```
CmdReturn(command string) (string, error)
    will run a command with either `bash -c` or `cmd /C` and return the output

CmdSTDOUT(command string)
    will run a command with either `bash -c` or `cmd /C` and print output and error to STDOUT

CmdNoOut(command string)
    will run a command with either `bash -c` or `cmd /C` and will provide no output

PSReturn(command string) (string, error)
    runs a command with `powershell -Command "& {}"` and returns the output as a string

PSSTDOUT(command string)
    runs a command with `powershell -Command "& {}"` and prints the output and error to STDOUT

PSNoOut(command string)
    runs a command with `powershell -Command "& {}"` and will provide no output
```

## File Manipulation
```
CopyFile(source string, destination string) error
    copy a file

MoveFile(source string, destination string) error
    move a file

Chmod(filepath string, permissions os.FileMode) error
    change the permissions of a file

ZipFiles(files []string, zipFileName string) error
    takes a slice of file paths and creates a zip archive. zipFileName should not include ".zip"

DoesFileExist(filepath string) bool
    check if a file exists. return true if it does
```