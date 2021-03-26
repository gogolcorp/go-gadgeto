package services
import "os"

// DirectoryExists returns true if directory exists, false elseway
func DirectoryExists(path string) bool {
	_, err := os.Stat(path)
	return os.IsExist(err)
}

// RemoveDirAndFiles removes the given directory and all its files
func RemoveDirAndFiles(path string) error {
	if err := os.RemoveAll(path); err != nil {
		return err
	}
	return os.Mkdir(path, os.ModePerm)
}