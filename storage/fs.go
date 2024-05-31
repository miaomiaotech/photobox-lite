package storage

import (
	"log"
	"os"
)

func CreateDirIfNotExist(dir string) error {
	mode := os.FileMode(0777)
	_, err := os.Stat(dir)
	if os.IsNotExist(err) {
		if err := os.MkdirAll(dir, mode); err != nil {
			return err
		}
	}
	return nil
}

// exist, permmission allowed -> true
// exist, permmission denied -> panic, because can't confirm exist or not
// not exist -> false
func Exist(path string) bool {
	_, err := os.Stat(path)
	// https://stefanxo.com/go-anti-patterns-os-isexisterr-os-isnotexisterr/
	// os.IsExist(err) does NOT work, because Stat will never return this error;
	// os.IsExist(err) is good for cases when you expect the file to not exist yet,
	// but the file actually exists, such as os.Symlink
	if err != nil {
		if os.IsNotExist(err) {
			return false
		} else {
			panic(err)
		}
	}
	return true
}

func PrepareDir(filePath string) {
	if err := os.MkdirAll(filePath, os.FileMode(0755)); err != nil {
		log.Fatal(err)
	}
}
