package hellsgopher

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"io"
	"os"
)

// return a random integer between min and max (both inclusive)
func RandomInt(min int, max int) int {
	return 0
}

// generate a random string of length length (Uses a-zA-Z)
func RandomStr(length int) string {
	return ""
}

// encode a string to base64 and return a string
func Base64EncodeStr(message string) string {
	encoded := base64.StdEncoding.EncodeToString([]byte(message))
	return encoded
}

// decode a string from base64 and return a string
func Base64DecodeStr(message string) (string, error) {
	decoded, err := base64.StdEncoding.DecodeString(message)
	if err != nil {
		return "", err
	}

	return string(decoded), nil
}

// get the md5 hash of a string
func MD5SumStr(message string) string {
	hash := md5.Sum([]byte(message))
	return hex.EncodeToString(hash[:])
}

// get the md5 hash of a file
func MD5SumFile(filepath string) string {
	if !DoesFileExist(filepath) {
		return ""
	}

	file, _ := os.Open(filepath)
	defer file.Close()

	hash := md5.New()
	io.Copy(hash, file)

	return hex.EncodeToString(hash.Sum(nil))
}

// get the sha1 hash of a string
func Sha1Str(message string) string {
	hash := sha1.Sum([]byte(message))
	return hex.EncodeToString(hash[:])
}

// get the sha1 hash of a file
func Sha1File(filepath string) string {
	if !DoesFileExist(filepath) {
		return ""
	}

	file, _ := os.Open(filepath)
	defer file.Close()

	hash := sha1.New()
	io.Copy(hash, file)

	return hex.EncodeToString(hash.Sum(nil))
}

// get the sha256 hash of a string
func Sha256Str(message string) string {
	hash := sha256.Sum256([]byte(message))
	return hex.EncodeToString(hash[:])
}

// get the sha256 hash of a file
func Sha256File(filepath string) string {
	if !DoesFileExist(filepath) {
		return ""
	}

	file, _ := os.Open(filepath)
	defer file.Close()

	hash := sha256.New()
	io.Copy(hash, file)

	return hex.EncodeToString(hash.Sum(nil))
}

// get the sha512 hash of a string
func Sha512Str(message string) string {
	hash := sha512.Sum512([]byte(message))
	return hex.EncodeToString(hash[:])
}

// get the sha512 hash of a file
func Sha512File(filepath string) string {
	if !DoesFileExist(filepath) {
		return ""
	}

	file, _ := os.Open(filepath)
	defer file.Close()

	hash := sha512.New()
	io.Copy(hash, file)

	return hex.EncodeToString(hash.Sum(nil))
}

// caesar cipher
func Caesar(message string, shift int) (string, error) {
	return "", ErrFuncNotSupported
}
