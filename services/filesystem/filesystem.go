package filesystem

import (
	"os"

	log "github.com/sirupsen/logrus"
)

// DirectoryExists returns true if directory exists, false elseway
func DirectoryExists(path string) bool {
	_, err := os.Stat(path)

	return !os.IsNotExist(err)
}

// RemoveDirAndFiles removes the given directory and all its files
func RemoveDirAndFiles(path string) error {
	if err := os.RemoveAll(path); err != nil {
		return err
	}
	return os.Mkdir(path, os.ModePerm)
}

// RemoveSingle removes a single resource from filesystem, file or directory
func RemoveSingle(path string) error {
	return os.Remove(path)
}

func GetWorkdirOrDie() string {
	workdir, err := os.Getwd()
	if err != nil {
		log.Fatal("Couldn't get workdir")
	}

	return workdir
}