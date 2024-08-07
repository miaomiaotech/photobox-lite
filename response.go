package main

import (
	"fmt"
	"log"
	"net/url"
	"path"
	"time"

	"github.com/miaomiaotech/photobox-lite/imageupload"
)

type UploadResponse struct {
	Image *imageupload.Image `json:"image"`
	Thumb *imageupload.Image `json:"thumb"`
	Data  *ImagePathUrl      `json:"data"`
}

type ImagePathUrl struct {
	OriginPath string `json:"path"`
	ThumbPath  string `json:"thumb_path"`
	OriginURL  string `json:"url"`
	ThumbURL   string `json:"thumb_url"`
}

func genImgPathUrl(hash, format string) ImagePathUrl {
	now := time.Now()
	if format == "jpeg" {
		format = "jpg"
	}

	suffix := path.Join(now.Format(dateFormat), fmt.Sprintf("%v.%v", hash, format))
	origin, thumb := path.Join("origin", suffix), path.Join("thumb", imageupload.ReplaceFileExt(suffix, "jpg"))
	return ImagePathUrl{
		OriginPath: origin,
		ThumbPath:  thumb,
		OriginURL:  joinURI(ImageDomain, origin),
		ThumbURL:   joinURI(ImageDomain, thumb),
	}
}

func joinURI(base, uri string) string {
	u, err := url.Parse(uri)
	if err != nil {
		log.Fatal(err)
	}
	b, err := url.Parse(base)
	if err != nil {
		log.Fatal(err)
	}
	return b.ResolveReference(u).String()
}
