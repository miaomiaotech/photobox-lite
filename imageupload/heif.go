package imageupload

import (
	_ "github.com/strukturag/libheif/go/heif"
)

func isHeifMime(contentType string) bool {
	return contentType == "image/heic" || contentType == "image/heif"
}
