package hellsgopher

// output a warning message to STDOUT
func Warn(message string) {
	print("[!] " + message + "\n")
}

// output error message to STDOUT
func Error(message string) {
	print("[-] " + message + "\n")
}

// output success mesage to STDOUT
func Okay(message string) {
	print("[+] " + message + "\n")
}

// output info message to STDOUT
func Info(message string) {
	print("[*] " + message + "\n")
}
