package main

import (
	"log"
	"net/http"
	"path"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/miaomiaotech/photobox-lite/imageupload"
	"github.com/miaomiaotech/photobox-lite/storage"
)

// APIUpload and generate thumbnail
func APIUpload(c *gin.Context) {
	img, err := imageupload.Process(c.Request, "file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// check cache
	cacheRes := UploadResponse{}
	err = CacheGet(c.Request.Context(), img.Md5, &cacheRes)
	if err == nil {
		if storage.Exist(cacheRes.Image.Path) && storage.Exist(cacheRes.Thumb.Path) {
			log.Printf("hit cache %v", img.Md5)
			c.JSON(http.StatusOK, cacheRes)
			return
		}
	}

	width, height, quality := getThumbParams(c)
	thumbnailImage, err := imageupload.Thumbnail(img, width, height, quality)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	pu := genImgPathUrl(img.Md5, img.Format)

	// save origin image
	err = saveImage(DataDir, pu.OriginPath, img)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// save thumbnail image
	err = saveImage(DataDir, pu.ThumbPath, thumbnailImage)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// save redis cache
	res := UploadResponse{img, thumbnailImage, &pu}
	err = CacheSet(img.Md5, &res)
	if err != nil {
		log.Println(err)
	}

	// response
	c.JSON(http.StatusOK, res)
}

func APIThumbnail(c *gin.Context) {
	img, err := imageupload.Process(c.Request, "file")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	width, height, quality := getThumbParams(c)
	t, err := imageupload.Thumbnail(img, width, height, quality)
	t.WriteResponse(c.Writer)
}

func saveImage(dir, keyPath string, img *imageupload.Image) error {
	fp := path.Join(dir, keyPath)
	local := storage.LocalStorage{Img: img}
	return storage.SaveTo(&local, fp)
}

func getThumbParams(c *gin.Context) (int, int, int) {
	width := defaultMaxWidth
	height := defaultMaxHeight
	quality := defaultQuality
	if str, ok := c.GetQuery("width"); ok {
		if v, e := strconv.Atoi(str); e == nil {
			width = v
		}
	}
	if str, ok := c.GetQuery("height"); ok {
		if v, e := strconv.Atoi(str); e == nil {
			height = v
		}
	}
	if str, ok := c.GetQuery("quality"); ok {
		if v, e := strconv.Atoi(str); e == nil {
			quality = v
		}
	}
	return width, height, quality
}
