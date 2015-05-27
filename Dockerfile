FROM golang:1.4.2
MAINTAINER Abdulkadir Yaman <abdulkadiryaman@gmail.com>

RUN mkdir /tmp/gopath
ENV GOPATH /tmp/gopath

RUN go get github.com/goksydome/colony

ENTRYPOINT ${GOPATH}/bin/colony
