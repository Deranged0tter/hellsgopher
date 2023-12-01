# hellsgopher
<img src="https://github.com/Deranged0tter/hellsgopher/blob/main/logo/logo.PNG?raw=true" data-canonical-src="https://gyazo.com/eb5c5741b6a9a16c692170a41a49c858.png" width="250" height="250" />
Go library for malware development

# ⚠️ THIS IS CURRENTLY UNDER DEVELOPMENT ⚠️

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
    runs a command with `powershell.exe` and returns the output as a string

PSSTDOUT(command string)
    runs a command with `powershell.exe` and prints the output and error to STDOUT

PSNoOut(command string)
    runs a command with `powershell.exe` and will provide no output
```

## File Manipulation
```
CopyFile(source string, destination string) error
    copy a file

MoveFile(source string, destination string) error
    move a file

DeleteFile(filepath string) error
    delete a file

Chmod(filepath string, permissions os.FileMode) error
    change the permissions of a file

ZipFiles(files []string, zipFileName string) error
    takes a slice of file paths and creates a zip archive. zipFileName should not include ".zip"

DoesFileExist(filepath string) bool
    check if a file exists. return true if it does

ListFiles(directory string) ([]string, error)
    returns a slice of files in a given directory

ListFilesInPWD() ([]string, error)
    returns a slice of files in the present working directory

ListFilesR(directory string) ([]string, error)
    returns a slice of files in a given directory recursivley

ListFilesInPWDR() ([]string, error)
    returns a slice of iles in the present working directory recursivley

DownloadFile(source string, destination string) error
    download a file from a source url to destination

ReadFileToSlice(filepath string) ([]string, error)
    read a file line by line and return a slice with each line as a value

ReadFileToStr(filepath string) (string, error)
    read a file and return a string of its content

ClearFile(filepath string)
    wipe the contents of a file

PrependToFile(filepath string, content string)
    prepend text to a file (will create a new first line)

AppendToFile(filepath string, content string)
    append text to a file (will create a new last line)
```

## Encryption
```
RandomInt(min int64, max int64) int64
    return a random integer between min and max

RandomStr(length int) string
    generates a random strength of length (uses a-zA-Z)

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

## System Enumeration
```
GetHostname() string
    returns the computer's hostname

GetDomainName() string
    return the domain name of the computer

GetOS() string
    return the operating system of the computer

GetOSBuild() string
    return the OS build of the computer

GetOSVersion() string
    return the OS version of the computer
```

## Logging Functions
```
Warn(message string)
    output a warning message to STDOUT ("[!] message")

Error(message string)
    output an error message to STDOUT ("[-] message")

Okay(message string)
    output a success message to STDOUT ("[+] message")

Info(message string)
    output an information message to STDOUT ("[*] message")
```

## Process Enumeration
```
ListProcesses() (map[int]string, error)
    returns a map of all process currently running

GetCurrentProcessID() int
    return the PID of the current process

GetCurrentParentProcessID() int
    return the PPID of the current process

GetCurrentProcessPath() (string, error)
    return the process path of the current process

GetCurrentProcessName() string
    return the name of the current process

GetCurrentProcessArch() string
    return the arch of the current process
```

## Network Enumeration
```
GetInternalIP() string
    get the internal IP of the computer

GetAllInterfaces() []string
    return a slice of all interfaces on computer
```

## Other
```
GenerateName() string
    return a string equaling to "adj_noun"
```
