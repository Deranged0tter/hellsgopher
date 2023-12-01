package hellsgopher

import (
	"archive/zip"
	"errors"
	"io"
	"os"
)

// copy a file
func CopyFile(source string, destination string) error {
	srcFile, err := os.Open(source)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	destFile, err := os.Create(destination)
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

// move a file
func MoveFile(source string, destination string) error {
	err := os.Rename(source, destination)
	return err
}

// delete a file
func DeleteFile(filepath string) error {
	return nil
}

// change the permissions of a file
func Chmod(filepath string, permissions os.FileMode) error {
	err := os.Chmod(filepath, permissions)
	return err
}

// takes a slice of file paths and creates a zip archive. zipFileName should not include ".zip"
func ZipFiles(files []string, zipFileName string) error {
	zipFile, err := os.Create(zipFileName + ".zip")
	if err != nil {
		return err
	}
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	for _, file := range files {
		toZip, err := os.Open(file)
		if err != nil {
			return err
		}
		defer toZip.Close()

	}

	return nil
}

// list files in a given directory
func ListFiles(directory string) ([]string, error) {
	return nil, nil
}

// list files in current directory
func ListFilesInPWD() ([]string, error) {
	return nil, nil
}

// list files in a given directory recursivley
func ListFilesR(directory string) ([]string, error) {
	return nil, nil
}

// list files in current directory recursivley
func ListFilesInPWDR() ([]string, error) {
	return nil, nil
}

// check if a file exists. return true if it does
func DoesFileExist(filepath string) bool {
	if _, err := os.Stat(filepath); errors.Is(err, os.ErrNotExist) {
		return false
	}

	return true
}

// download a file from source url to destination. if destination is left blank, will default to current directory with same filename from url
func DownloadFile(source string, destination string) error {
	return nil
}

// read a file line by line, putting each line into a slice
func ReadFileToSlice(filepath string) ([]string, error) {
	return nil, nil
}

// read a file and return a string of its content
func ReadFileToStr(filepath string) (string, error) {
	return "", nil
}

// clears the content of a file
func ClearFile(filepath string) {

}

// prepends a string to a file
func PrependToFile(filepath string, content string) {

}

// append a string to a file
func AppendToFile(filepath string, content string) {

}
