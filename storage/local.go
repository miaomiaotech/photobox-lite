package storage

import (
	"os"
	"path"

	"github.com/miaomiaotech/photobox-lite/imageupload"
)

type LocalStorage struct {
	Img *imageupload.Image
}

func (s *LocalStorage) Save(fp string) error {
	err := CreateDirIfNotExist(path.Dir(fp))
	if err != nil {
		return err
	}

	return s.Img.Save(fp)
}
func (s *LocalStorage) Read(fp string) ([]byte, error) {
	return os.ReadFile(fp)
}
