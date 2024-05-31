FROM golang:1.22.3
WORKDIR /app
ENV DOMAIN=https://mem.drink.cafe
VOLUME /images

RUN apt update; apt install -y curl htop vim redis tmux bash

# https://github.com/MaestroError/heif-converter-image/blob/maestro/install-libheif.sh
RUN apt-get install -y git cmake make pkg-config libx265-dev libde265-dev libjpeg-dev libtool && \
    git clone https://github.com/strukturag/libheif.git --depth 1
RUN cd libheif && mkdir build && cd build && cmake .. && make && make install && ldconfig && echo "Installation of libheif is completed." && rm -rf libheif

COPY . .
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go build -ldflags "-s -w" -o /bin/app .

ENV LISTEN=":8000"
CMD redis-server --daemonize yes && app -listen $LISTEN -domain $DOMAIN
