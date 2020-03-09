#!/bin/sh

#wget -q http://geolite.maxmind.com/download/geoip/database/GeoLite2-Country.tar.gz

wget  -q "https://download.maxmind.com/app/geoip_download?edition_id=GeoLite2-Country&license_key=nRVkwS5Q0bWQTyAq&suffix=tar.gz" -O  GeoLite2-Country.tar.gz

tar zxf GeoLite2-Country.tar.gz

#ONCE if you miss the tool: 
## go get -u github.com/jteeuwen/go-bindata/...
go-bindata -o db.go GeoLite2-Country_*/GeoLite2-Country.mmdb

grep -A1 sources db.go | tail -1 | awk '{ print "package main\n// Assetname is the Asset name\nvar Assetname=£"$2"£"  }'| sed "s/£/\"/g" > name.go
