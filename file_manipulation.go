package hellsgopher

import (
	"archive/zip"
	"bufio"
	"errors"
	"io"
	"io/fs"
	"net/http"
	"os"
	"path/filepath"
)

// copy a file from sourcePath to destinationPath
func CopyFile(sourcePath string, destinationPath string) error {
	srcFile, err := os.Open(sourcePath)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	destFile, err := os.Create(destinationPath)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, srcFile)
	if err != nil {
		return err
	}

	err = destFile.Sync()
	if err != nil {
		return err
	}

	return nil
}

// move a file from sourcePath to destinationPath
func MoveFile(sourcePath string, destinationPath string) error {
	err := os.Rename(sourcePath, destinationPath)
	return err
}

// delete a file
func DeleteFile(path string) error {
	err := os.Remove(path)
	return err
}

// delete a directory
func DeleteDir(dir string) error {
	err := os.RemoveAll(dir)
	return err
}

// change permissions of a file
func Chmod(path string, perms os.FileMode) error {
	err := os.Chmod(path, perms)
	return err
}

// take a slice of file paths and creates a zip archive
// note: zipFileName should not include ".zip"
func ZipFiles(paths []string, zipFileName string) error {
	zipFile, err := os.Create(zipFileName + ".zip")
	if err != nil {
		return err
	}
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	for _, file := range paths {
		toZip, err := os.Open(file)
		if err != nil {
			return err
		}
		defer toZip.Close()

	}

	return nil
}

// check if a file exists
// returns true if file exists
func DoesFileExist(path string) bool {
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		return false
	}

	return true
}

// return the present working dir
func GetPwd() (string, error) {
	pwd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	return pwd, nil
}

// returns a slice of files in a given dir
func ListFiles(dir string) ([]string, error) {
	var files []string

	err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		files = append(files, path)
		return nil
	})

	return files, err
}

// return a slice of files in the present working dir
func ListFilesInPwd() ([]string, error) {
	var files []string
	pwd, err := GetPwd()
	if err != nil {
		return files, err
	}

	err = filepath.WalkDir(pwd, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		files = append(files, path)
		return nil
	})

	return files, err
}

// download a file from a source url to a destination path
func DownFile(source string, dest string) error {
	// create file
	file, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer file.Close()

	// create request
	req, err := http.NewRequest("GET", source, nil)
	if err != nil {
		return err
	}

	// set headers to blend in a bit better
	req.Header.Set("Accept", "application/x-www-form-urlencoded")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/106.0.0.0 Safari/537.36")

	// do request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// write body of response to file
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

// read a file line by line and return a slice with each line as a value
func ReadFileToSlice(path string) ([]string, error) {
	var lines []string

	file, err := os.Open(path)
	if err != nil {
		return lines, err
	}

	fScanner := bufio.NewScanner(file)
	fScanner.Split(bufio.ScanLines)

	for fScanner.Scan() {
		lines = append(lines, fScanner.Text())
	}
	file.Close()

	return lines, nil
}

// read a file and return a string of its content
func ReadFileToString(path string) (string, error) {
	var content string

	lines, err := ReadFileToSlice(path)
	if err != nil {
		return "", err
	}

	for _, line := range lines {
		content += line + "\n"
	}

	return content, nil
}

// wipe a file of all its contents
func WipeFile(path string) error {
	err := os.Truncate(path, 0)
	return err
}

// prepend text to a file
// creates a new first line
func PrependToFile(path string, s string) error {
	fContent, err := ReadFileToString(path)
	if err != nil {
		return err
	}

	newContent := s + "\n" + fContent

	file, err := os.Open(path)
	if err != nil {
		return err
	}

	_, err = file.Write([]byte(newContent))
	if err != nil {
		return err
	}
	defer file.Close()

	return nil
}

// append text to a file
// creates a new last line
func AppendToFile(path string, s string) error {
	fContent, err := ReadFileToString(path)
	if err != nil {
		return err
	}

	newContent := fContent + "\n" + s

	file, err := os.Open(path)
	if err != nil {
		return err
	}

	_, err = file.Write([]byte(newContent))
	if err != nil {
		return err
	}
	defer file.Close()

	return nil
}

// create a new blank file
func NewFile(path string) error {
	_, err := os.Create(path)
	return err
}

// create a new file containing content
func NewFileWithContent(path string, content string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	file.Write([]byte(content))

	return nil
}
