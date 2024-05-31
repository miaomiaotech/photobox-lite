package imageupload

import (
	"log"

	"github.com/klippa-app/go-libheif"
	"github.com/klippa-app/go-libheif/library"
)

// see https://github.com/klippa-app/go-libheif
func heicContentType(contentType string) bool {
	return contentType == "image/heic" || contentType == "image/heif"
}

// go build ./cmd/libheif-plugin
func InitLibHeifWorker(workerBinPath string) {
	err := libheif.Init(libheif.Config{LibraryConfig: library.Config{
		Command: library.Command{BinPath: workerBinPath},
	}})
	if err != nil {
		log.Fatalf("could not start libheif worker: %s", err.Error())
	}
}
