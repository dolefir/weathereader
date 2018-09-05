FROM golang:alpine
LABEL maintainer="Olefir Dmytro <itfortim@gmail.com>"

EXPOSE 5050
RUN apk --update-cache --allow-untrusted \
        --repository http://dl-4.alpinelinux.org/alpine/edge/community \
        --arch=x86_64 add \
    glide make git \
    && rm -rf /var/cache/apk/* \
    go get github.com/pilu/fresh 


COPY . $GOPATH/src/github.com/simplewayUA/weathereader

WORKDIR $GOPATH/src/github.com/simplewayUA/weathereader

CMD  go run main.go
