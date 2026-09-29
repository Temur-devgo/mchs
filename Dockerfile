FROM golang:1.12.0-alpine3.9

RUN mkdir /go/src/mchs-api

ADD . /go/src/mchs-api

WORKDIR /go/src/mchs-api

RUN sed -i 's/dl-cdn.alpinelinux.org/mirror.neolabs.kz/g' /etc/apk/repositories \
    && apk --no-cache add git

RUN go get

CMD ["go", "run", "main.go"]