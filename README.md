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

## API

```
/upload
/thumbnail

    Common query parameters:
        width     int    "max thumbnail width"
        height    int    "max thumbnail width"
        quality   int    "thumbnail quality"
```