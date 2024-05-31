package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"strings"

	"github.com/gin-gonic/gin"
)

var (
	DataDir          = "photos"
	ImageDomain      = "http://photobox.bitsflow.org"
	listen           = ":5000"
	redisDefault     = "127.0.0.1:6379"
	redisDB          = 8
	defaultMaxWidth  = 1280
	defaultMaxHeight = 1280
	defaultQuality   = 80
)

func init() {
	flag.StringVar(&DataDir, "data", DataDir, "directory to save photos")
	flag.StringVar(&ImageDomain, "domain", ImageDomain, "photos storage domain name")
	flag.StringVar(&listen, "listen", listen, "bind [<host>]:<port>")
	flag.StringVar(&redisDefault, "redis", redisDefault, "redis <host>:<port>")
	flag.IntVar(&redisDB, "db", redisDB, "redis database")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Please setup AWS_ACCESS_KEY_ID, AWS_SECRET_ACCESS_KEY, AWS_DEFAULT_REGION, PHOTOBOX_BUCKET(default photobox-develop) to use S3 as L2 storage.\n\n")
		flag.PrintDefaults()
	}
	flag.Parse()
	if !strings.HasPrefix(ImageDomain, "http") {
		ImageDomain = "http://" + ImageDomain
	}

	initCodec()
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
