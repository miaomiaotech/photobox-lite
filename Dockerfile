FROM golang:1.22.3
WORKDIR /app
ENV DOMAIN=https://mem.drink.cafe
VOLUME /photos

RUN apt update; apt install -y curl htop vim redis tmux bash

# https://github.com/MaestroError/heif-converter-image/blob/maestro/install-libheif.sh
RUN apt-get install -y git cmake make pkg-config libx265-dev libde265-dev libjpeg-dev libtool && \
    git clone https://github.com/strukturag/libheif.git --depth 1
RUN cd libheif && mkdir build && cd build && cmake .. && make && make install && ldconfig && echo "Installation of libheif is completed."

COPY . .
RUN go env -w GOPROXY=https://goproxy.cn,direct; go build -ldflags "-s -w" -o /bin/app . && cp /app/index.html / && rm -rf /app/* && mv /index.html /app

ENV LISTEN=":8000"
CMD redis-server --daemonize yes && app -listen $LISTEN -domain $DOMAIN -data /photos -callback after-upload.sh
