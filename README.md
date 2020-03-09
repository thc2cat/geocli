# GeoCli

## Purpose
geocli is a oneline CLI that provide geoip and dns checkout.
geocli come with geolite country database directly embedded into the cli binary.

## Needs
 - Create an account on Maxmind site for downloading database.

## Build
 - Makefile download maxmind geolite database.
 - Makefile build name.go from assets with go-bindata.
 - Golang then produce the binary 

## Usage :
  - "-v" is used to show maxmind geolite database version.
  - "-r" option read stdin and try to resolve dns in parrallel.
  - otherwise arg1 is checked.

## Exemple :
 
 ```
 $/local/bin/geocli -v 
/local/bin/geocli is using GeoLite2-Country_20191001/GeoLite2-Country.mmdb

 lune# rg "imap\[.* login: " /var/log/imapd.log | cut -d: -f5- | sed "s/ User.*//g" | rg -v "webmail|TLS" | sort -u|rg -v uvsq | cut -d\[ -f2 | cut -d\] -f1 | sor
37.171.85.185 [37-171-85-185.coucou-networks.fr.] France, FR
10.172.16.103 [unknown] , 
37.164.163.41 [unknown] France, FR
37.170.73.3 [37-170-73-3.coucou-networks.fr.] France, FR
37.164.246.169 [unknown] France, FR
176.179.79.3 [176-179-79-3.abo.bbox.fr.] France, FR
37.170.3.114 [37-170-3-114.coucou-networks.fr.] France, FR
37.171.67.201 [37-171-67-201.coucou-networks.fr.] France, FR
```

 ## References :
  - go-bindata : 
  > go get -u github.com/jteeuwen/go-bindata/...
