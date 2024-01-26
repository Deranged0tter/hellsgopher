# hellsgopher
<img src="https://github.com/Deranged0tter/hellsgopher/blob/main/logo/logo.PNG?raw=true" data-canonical-src="https://gyazo.com/eb5c5741b6a9a16c692170a41a49c858.png" width="250" height="250" />
Go library for malware development

# To add to your project
```
go get github.com/deranged0tter/hellsgopher
```

I recommend importing hellsgopher in the following manner:
```
import (
    hg "github.com/deranged0tter/hellsgopher"
)
```
This way, you can use hg.FUNCTION()

# Goal
The goal of hellsgopher is to make malware development easier to learn and get into. Since this source is public, it will most likely get burned and caught by AVs. As such it is not intended for actual use, but for learning purposes. This library is designed to only work on windows.

# Contributing
see [contributing](CONTRIBUTING.md)

# Functions
## Command Line Functions
```
CmdReturn(command string) (string, error)
    will run cmd.exe and return output

CmdStdOUT(command string)
    will run cmd.exe and print output to STDOUT

CmdNoOut(command string)
    will run cmd.exe and provide no output

PsReturn(command string) (string, error)
    will run powershell command and return output

PsStdOut(command string)
    will run powershell command and print output to STDOUT

PsNoOut(command string)
    will run powershell command and provide no output

PsReturnT(command string, token windows.Token) (string, error)
    will run powershell command and return output (with token)

PsStdOutT(command string, token windows.Token)
    will run powershell command and print output to STDOUT (with token)

PsNoOutT(command string, token windows.Token)
    will run powershell command and provide no output (with token)
```

## File Manipulation Functions
```
CopyFile(sourcePath string, destinationPath string) error
    copy a file from sourcePath to destinationPath

MoveFile(sourcePath string, destinationPath string) error
    move a file from sourcePath to destinationPath

DeleteFile(path string) error
    delete a file

DeleteDir(dir string) error
    delete a directory

Chmod(path string, perms os.FileMode) error
    change permissions of a file

ZipFiles(paths []string, zipFileName string) error
    take a slice of file paths and creates a zip archive
    note: zipFileName should not include ".zip"

DoesFileExist(path string) bool
    check if a file exists
    returns true if file exists

GetPwd() (string, error)
    return the present working dir

ListFiles(dir string) ([]string, error)
    returns a slice of files in a given dir

ListFilesInPwd() ([]string, error)
    return a slice of files in the present working dir

DownFile(source string, dest string) error
    download a file from a source url to a destination path

ReadFileToSlice(path string) ([]string, error)
    read a file line by line and return a slice with each line as a value

ReadFileToString(path string) (string, error)
    read a file and return a string of its content

WipeFile(path string) error
    wipe a file of all its contents (truncates the file)

PrependToFile(path string, s string) error
    prepend text to a file
    creates a new first line

AppendToFile(path string, s string) error
    append text to a file
    creates a new last line

NewFile(path string) error
    create a new blank file

NewFileWithContent(path string, content string) error
    create a new file containing content
```

## Encryption Functions
```
RandomInt(min int, max int) (int, error)
    return a random int between min and max

RandomStr(l int) string
    return a random string of length l
    uses a-zA-Z

RandomStrI(l int) string
    returns a random string combining letters and numbers of length l
    uses a-zA-Z0-9

RandomStrFromCharset(l int, charset string) string
    returns a random string from provided charset of length l

Base64EncodeStr(s string) string
    encode a string to base64

Base64DecodeStr(s string) (string, error)
    decode a string from base64

Base32EncodeStr(s string) string
    encode a string to 
    
Base32DecodeStr(s string) (string, error)
    decode a string from base32

Md5String(s string) string
    get the md5 hash of a string

Md5File(path string) string
    get the md5 hash of a file

Sha1String(s string) string
    get the sha1 hash of a string

Sha1File(path string) string
    get the sha1 hash of a file

Sha256String(s string) string
    get the sha256 hash of a string

Sha256File(path string) string
    get the sha256 hash of a file

Sha512String(s string) string
    get the sha512 hash of a string

Sha512File(path string) string
    get the sha512 hash of a file

RotX(s string, shift rune) string
    rot cipher
```

## User Enumeartion Functions
```
GetCurrentUser() (*user.User, error)
    return a user.User for the current user

GetCurrentUsername() (string, error)
    get the current username

GetCurrentUid() (string, error)
    get the current uid

GetCurrentGid() (string, error)
    get the main gid for the current user

GetCurrentGids() ([]string, error)
    get all gids for the current user

GetUidFromName(name string) (string, error)
    return a uid from a given username

GetNameFromUid(uid string) (string, error)
    return a username from a given uid

GetUserFromName(name string) (*user.User, error)
    return a user.User from username

GetUserFromUid(uid string) (*user.User, error)
    return a user.User from uid

GetAllUsers() ([]*user.User, error)
    return a slice of all users on the machine

GetAllUsernames() ([]string, error)
    return a slice of all usernames on the machine
```

## System Enumeration Functions
```
GetHostname() (string, error)
    return the machine's hostname

GetDomainName() (string, error)
    return the domain name of the machine

GetOS() string
    return the machine's OS

GetOSBuild() string
    return the machine's OS Build Number

GetOSVersion() string
    return the machine's OS Version
```

## Process Enumeration Functions
```
GetPidFromName(name string) ([]int, error)
    return the pid(s) from the process name

GetNameFromPid(pid int) (string, error)
    get the name from the pid

ListAllProcesses() ([]WinProcess, error)
    list all running processes

GetCurrentPid() int
    get the pid of current process

GetCurrentPpid() int
    get ppid of current process

GetCurrentProcPath() (string, error)
    get the path of the current process

GetCurrentProcName() (string, error)
    get the name of the current process

GetCurrentProcArch() string
    get the arch of the current process
```

## Network Enumeration Functions

## Shellcode Functions

## Injection Functions

## Scanning Functions

## Evasion Functions

## Anti-Sandbox Functions
If you only want the Anti-Sandboxing functions, I have a [library](https://github.com/deranged0tter/govm) for you!

## Anti-Forensics Functions

## Token Manipulation Functions
```
GetCurrentToken() (windows.Token, error)
    get the token from the current process

GetTokenFromPid(pid int) (windows.Token, error)
    get the token from a process given its pid

GetTokenFromName(procName string) (windows.Token, error)
    get the token from a process given its process name
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

## Other Functions

# 3rd Party Libraries
```
github.com/fourcorelabs/wintoken
```
Thank you to the creators of these projects!

# Liability
The creator nor any person who has contributed to this project is liable for any kind of malicious of illegal use of this software. Only use this on targets, systems, networks, etc that you have own and/or have permission to use on.

DO NOT USE THIS FOR:
- illegal actions
- malicious actions
- damaging actions to property you do not have direct permission to use this on

Any use of this software for illegal actions is not the responsibility of the creator or any contributor of this project. We hold no liability for any actions taken by this software.