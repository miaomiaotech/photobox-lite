# Photo Box Lite

![logo](camera-logo.png)

## Features

- API
    - `/upload`
    - `/thumbnail`
- Save origin image on local disk
- Generate thumbnail image with size you desire
- JSON response format
- Custom `index.html` homepage
- Redis cached upload result based on photo hash

## Try in local environment

```bash
brew install libheif # see https://github.com/klippa-app/go-libheif for linux
go build -ldflags "-s -w" ./cmd/libheif-plugin/
LIB_HEIF_PLUGIN=$PWD/libheif-plugin go run *.go -listen '127.0.0.1:8000' -domain 'http://127.0.0.1:8000'
```

## API

```
/upload
/thumbnail

    Common query parameters:
        width     int    "max thumbnail width"
        height    int    "max thumbnail width"
        quality   int    "thumbnail quality"
```