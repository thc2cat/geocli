#include  ../../make/Makefile-for-go.mk

all: | name.go 
	go build .

name.go:
	sh get_extract_give_name.sh

clean:
	 @go clean
	 @rm -fr GeoLite2-Country* db.go name.go

