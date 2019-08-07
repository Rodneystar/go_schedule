FROM golang:1.12-alpine

RUN apk add git
RUN go get -u github.com/c9s/gomon
