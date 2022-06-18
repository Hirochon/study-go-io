FROM golang:1.18.2-alpine3.15

WORKDIR /go/src/work
COPY . /go/src/work/

RUN apk update && \
    apk add --no-cache git gcc musl-dev

EXPOSE 6868
