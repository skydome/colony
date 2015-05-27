FROM golang:latest
MAINTAINER Abdulkadir Yaman <abdulkadiryaman@gmail.com>

RUN mkdir /tmp/gopath
ENV GOPATH /tmp/gopath

RUN go get github.com/goskydome/colony

ENTRYPOINT ${GOPATH}/bin/colony
