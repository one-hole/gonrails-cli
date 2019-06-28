package helper

import "os"

// FileExists - file exists or not
func FileExists(path string) bool {
	_, err := os.Stat(path)
	if nil == err {
		return true
	}
	return false
}

// CreateOrOpenFile -
func CreateOrOpenFile(path string) (*os.File, error) {

	var file *os.File
	var err error

	if file, err = os.OpenFile(path, os.O_RDWR, os.ModeAppend); err != nil {
		file, err = os.Create(path)
	}

	return file, err
}
