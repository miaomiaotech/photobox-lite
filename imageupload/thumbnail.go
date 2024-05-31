package imageupload

import (
	"bytes"
	"image"
	"image/jpeg"
	"image/png"
	"path"
	"strings"

	"github.com/nfnt/resize"
)

// Create JPEG thumbnail.
func ThumbnailJPEG(i *Image, width int, height int, quality int) (*Image, error) {
	img, format, err := image.Decode(bytes.NewReader(i.Data))
	thumbnail := resize.Thumbnail(uint(width), uint(height), img, resize.Lanczos3)

	buf := new(bytes.Buffer)
	err = jpeg.Encode(buf, thumbnail, &jpeg.Options{
		Quality: quality,
	})
	if err != nil {
		return nil, err
	}

	data := buf.Bytes()
	t := &Image{
		Filename:    ReplaceFileExt(path.Base(i.Filename), "jpg"),
		ContentType: "image/jpeg",
		Format:      format,
		Data:        data,
		Size:        len(data),
		Width:       thumbnail.Bounds().Max.X,
		Height:      thumbnail.Bounds().Max.Y,
		Sha256:      Sha256(data),
		Md5:         Md5(data),
	}
	return t, nil
}

// Create PNG thumbnail.
func ThumbnailPNG(i *Image, width int, height int) (*Image, error) {
	img, format, err := image.Decode(bytes.NewReader(i.Data))
	thumbnail := resize.Thumbnail(uint(width), uint(height), img, resize.Lanczos3)

	buf := new(bytes.Buffer)
	err = png.Encode(buf, thumbnail)
	if err != nil {
		return nil, err
	}

	data := buf.Bytes()
	t := &Image{
		Filename:    ReplaceFileExt(path.Base(i.Filename), "png"),
		ContentType: "image/png",
		Format:      format,
		Data:        data,
		Size:        len(data),
		Width:       thumbnail.Bounds().Max.X,
		Height:      thumbnail.Bounds().Max.Y,
		Sha256:      Sha256(data),
		Md5:         Md5(data),
	}
	return t, nil
}

func Thumbnail(img *Image, width, height, quality int) (*Image, error) {
	if strings.HasSuffix(strings.ToLower(img.Filename), ".png") {
		return ThumbnailPNG(img, width, height)
	} else {
		return ThumbnailJPEG(img, width, height, quality)
	}
}

func ReplaceFileExt(filePath, newExtension string) string {
	if newExtension != "" && newExtension[0] != '.' {
		newExtension = "." + newExtension
	}

	// 如果文件路径中不包含点，则直接添加新后缀名
	if strings.LastIndex(filePath, ".") == -1 {
		return filePath + newExtension
	}

	fileName := filePath[:strings.LastIndex(filePath, ".")]
	newExtension = strings.ToLower(newExtension)
	return fileName + newExtension
}
