package main

import (
	"flag"
	"log"
	"os"
	"path"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/miaomiaotech/photobox-lite/imageupload"
)

var (
	DataDir               = "photos"
	ImageDomain           = "http://photobox.drink.cafe"
	listen                = ":5000"
	redisDefault          = "127.0.0.1:6379"
	redisDB               = 8
	defaultThumbMaxWidth  = 1280
	defaultThumbMaxHeight = 1280
	defaultQuality        = 80
)

func init() {
	flag.StringVar(&DataDir, "data", DataDir, "directory to save photos")
	flag.StringVar(&ImageDomain, "domain", ImageDomain, "photos storage domain name")
	flag.StringVar(&listen, "listen", listen, "bind [<host>]:<port>")
	flag.StringVar(&redisDefault, "redis", redisDefault, "redis <host>:<port>")
	flag.IntVar(&redisDB, "db", redisDB, "redis database")
	flag.Parse()
	if !strings.HasPrefix(ImageDomain, "http") {
		ImageDomain = "http://" + ImageDomain
	}

	initRedisCache()
	imageupload.InitLibHeifWorker(os.Getenv("LIB_HEIF_PLUGIN"))
}

func main() {
	r := gin.Default()
	r.StaticFS("/origin", gin.Dir(path.Join(DataDir, "origin"), false))
	r.StaticFS("/thumb", gin.Dir(path.Join(DataDir, "thumb"), false))

	r.GET("/", func(c *gin.Context) {
		c.File("index.html")
	})

	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"success": false, "message": "resource not found"})
	})

	r.POST("/upload", APIUpload)
	r.POST("/thumbnail", APIThumbnail)

	log.Fatal(r.Run(listen))
}
