#include  ../../make/Makefile-for-go.mk

NAME= $(notdir $(shell pwd))
TAG=$(shell git tag)


build: | init.go
	@go build -ldflags '-w -s -X main.Version=${NAME}-${TAG}'
	@notify-send 'Build Complete' 'Your project has been build successfully!' -u normal -t 7500 -i checkbox-checked-symbolic


init.go:
	sh get_extract_give_name.sh

clean:
	 @go clean
	 @rm -fr GeoLite2-Country* init.go  geocli*.bsd

bsd:
	 GOOS=freebsd GOARCH=amd64 go build -ldflags '-w -s -X main.Version=${NAME}-${TAG}' -o ${NAME}-${TAG}.bsd
