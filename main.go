package main

import (
	"flag"
	"log"
	"os"
	"path"
	"strings"

	"github.com/gin-gonic/gin"
)

var (
	DataDir      = "photos"
	ImageDomain  = "http://photobox.drink.cafe"
	listen       = ":5000"
	redisDefault = "127.0.0.1:6379"
	redisDB      = 8

	dateFormat            = "2006-01" // one directory each month
	maxWidthOrHeight      = 4000
	maxQuality            = 98
	defaultThumbMaxWidth  = 1280
	defaultThumbMaxHeight = 1280
	defaultQuality        = 95
	uploadCallback        = ""
	PASSWORD              = ""
	TEST_CALLBACK         = os.Getenv("TEST_CALLBACK") == "1"
)

func init() {
	flag.StringVar(&DataDir, "data", DataDir, "directory to save photos")
	flag.StringVar(&ImageDomain, "domain", ImageDomain, "photos storage domain name")
	flag.StringVar(&listen, "listen", listen, "bind [<host>]:<port>")
	flag.StringVar(&redisDefault, "redis", redisDefault, "redis <host>:<port>")
	flag.StringVar(&uploadCallback, "callback", uploadCallback, "executable to run after upload a image")
	flag.StringVar(&PASSWORD, "password", PASSWORD, "password to upload")
	flag.IntVar(&redisDB, "db", redisDB, "redis database")
	flag.Parse()
	if !strings.HasPrefix(ImageDomain, "http") {
		ImageDomain = "http://" + ImageDomain
	}
}

func init2() {
	initRedisCache()
}

func main() {
	if TEST_CALLBACK {
		code, stdout, stderr := RunSimpleCommand(uploadCallback)
		log.Printf("exit code: %d", code)
		log.Printf("stdout: %s", stdout)
		log.Printf("stderr: %s", stderr)
		return
	}

	init2()

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
