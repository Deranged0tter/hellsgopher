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

## Encryption
```
Base64EncodeStr(message string) string
    encode a string to base64 and return a string

Base64DecodeStr(message string) (string, error)
    decode a string from base64 and return a string

MD5SumStr(message string) string
    get the md5 hash of a string

MD5SumFile(filepath string) string
    get the md5 hash of a file

Sha1Str(message string) string
    get the sha1 hash of a string

Sha1File(filepath string) string
    get the sha1 hash of a file

Sha256Str(message string) string
    get the sha256 hash of a string

Sha256File(filepath string) string
    get the sha256 hash of a file

Sha512Str(message string) string
    get the sha512 hash of a string

Sha512File(filepath string) string
    get the sha512 hash of a file

Caesar(message string, shift int) (string, error)
    caesar cipher
```

## User Enumeration
```
GetUsername() (string, error)
    returns the username of the user the process is running under

GetDisplayName() (string, error)
    returns the display name of the user the process is running under

GetUID() (string, error)
    returns the user id of the user the process is running under

GetGID() (string, error)
    returns the group id of the user the process is running under

ListUsers() ([]string, error)
    list all the users on the systen
```