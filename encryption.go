package hellsgopher

import (
	"crypto/md5"
	cr "crypto/rand"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base32"
	"encoding/base64"
	"encoding/hex"
	"io"
	"math/big"
	"math/rand"
	"os"
	"strings"
)

// return a random int between min and max
func RandomInt(min int, max int) (int, error) {
	bg := big.NewInt(int64(max) - int64(min))

	n, err := cr.Int(cr.Reader, bg)
	if err != nil {
		return 0, err
	}

	return int(n.Int64()) + min, nil
}

// return a random string of length l
// uses a-zA-Z
func RandomStr(l int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	sb := strings.Builder{}
	sb.Grow(l)
	for i := 0; i < l; i++ {
		sb.WriteByte(charset[rand.Intn(len(charset))])
	}
	return sb.String()
}

// returns a random string combining letters and numbers of length l
// uses a-zA-Z0-9
func RandomStrI(l int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	sb := strings.Builder{}
	sb.Grow(l)
	for i := 0; i < l; i++ {
		sb.WriteByte(charset[rand.Intn(len(charset))])
	}
	return sb.String()
}

// returns a random string from provided charset of length l
func RandomStrFromCharset(l int, charset string) string {
	sb := strings.Builder{}
	sb.Grow(l)
	for i := 0; i < l; i++ {
		sb.WriteByte(charset[rand.Intn(len(charset))])
	}
	return sb.String()
}

// encode a string to base64
func Base64EncodeStr(s string) string {
	encoded := base64.StdEncoding.EncodeToString([]byte(s))
	return encoded
}

// decode a string from base64
func Base64DecodeStr(s string) (string, error) {
	decoded, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return "", err
	}

	return string(decoded), nil
}

// encode a string to base32
func Base32EncodeStr(s string) string {
	encoded := base32.StdEncoding.EncodeToString([]byte(s))
	return encoded
}

// decode a string from base32
func Base32DecodeStr(s string) (string, error) {
	decoded, err := base32.StdEncoding.DecodeString(s)
	if err != nil {
		return "", err
	}

	return string(decoded), nil
}

// get the md5 hash of a string
func Md5String(s string) string {
	hash := md5.Sum([]byte(s))
	return hex.EncodeToString(hash[:])
}

// get the md5 hash of a file
func Md5File(path string) string {
	if !DoesFileExist(path) {
		return ""
	}

	file, _ := os.Open(path)
	defer file.Close()

	hash := md5.New()
	io.Copy(hash, file)

	return hex.EncodeToString(hash.Sum(nil))
}

// get the sha1 hash of a string
func Sha1String(s string) string {
	hash := sha1.Sum([]byte(s))
	return hex.EncodeToString(hash[:])
}

// get the sha1 hash of a file
func Sha1File(path string) string {
	if !DoesFileExist(path) {
		return ""
	}

	file, _ := os.Open(path)
	defer file.Close()

	hash := sha1.New()
	io.Copy(hash, file)

	return hex.EncodeToString(hash.Sum(nil))
}

// get the sha256 hash of a string
func Sha256String(s string) string {
	hash := sha256.Sum256([]byte(s))
	return hex.EncodeToString(hash[:])
}

// get the sha256 hash of a file
func Sha256File(path string) string {
	if !DoesFileExist(path) {
		return ""
	}

	file, _ := os.Open(path)
	defer file.Close()

	hash := sha256.New()
	io.Copy(hash, file)

	return hex.EncodeToString(hash.Sum(nil))
}

// get the sha512 hash of a string
func Sha512String(s string) string {
	hash := sha512.Sum512([]byte(s))
	return hex.EncodeToString(hash[:])
}

// get the sha512 hash of a file
func Sha512File(path string) string {
	if !DoesFileExist(path) {
		return ""
	}

	file, _ := os.Open(path)
	defer file.Close()

	hash := sha512.New()
	io.Copy(hash, file)

	return hex.EncodeToString(hash.Sum(nil))
}

// rot cipher
func RotX(s string, shift rune) string {
	result := make([]string, 0, len(s))
	for _, chr := range s {
		if 'a' <= chr && chr <= 'z' {
			chr = ((chr - 'a' + shift) % 26) + 'a'
		}
		if 'A' <= chr && chr <= 'Z' {
			chr = ((chr - 'A' + shift) % 26) + 'A'
		}
		result = append(result, string(chr))
	}
	output := strings.Join(result, "")
	return output
}
