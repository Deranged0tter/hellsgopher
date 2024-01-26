package hellsgopher

// output "[!] message" to STDOUT
func Warn(message string) {
	print("[!] " + message + "\n")
}

// output "[-] error" to STDOUT
func Error(err string) {
	print("[-] " + err + "\n")
}

// output "[+] message" to STDOUT
func Okay(message string) {
	print("[+] " + message + "\n")
}

// output "[*] message" to STDOUT
func Info(message string) {
	print("[*] " + message + "\n")
}
