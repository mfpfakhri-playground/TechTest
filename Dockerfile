FROM golang:1.16.2

RUN mkdir /build
WORKDIR /build

RUN export GO111MODULE=on
RUN go get github.com