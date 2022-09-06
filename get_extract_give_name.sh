#!/bin/sh

. ./license_key

#wget -q "https://download.maxmind.com/app/geoip_download?edition_id=GeoLite2-Country&license_key=nRVkwS5Q0bWQTyAq&suffix=tar.gz" -O  GeoLite2-Country.tar.gz
wget -q "https://download.maxmind.com/app/geoip_download?edition_id=GeoLite2-Country&license_key=$license_key&suffix=tar.gz" -O  GeoLite2-Country.tar.gz

tar zxf GeoLite2-Country.tar.gz

GEOIPDATANAME=`ls GeoLite2-Country_*/GeoLite2-Country.mmdb`

sed "s+%%GEOIPDATANAME%%+$GEOIPDATANAME+g" < template/init_go.tpl > init.go
