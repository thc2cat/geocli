# GeoCli

## Purpose

`geocli` is a CLI that provide Geoip and dns checkout.
`geocli` come with geolite country database directly embedded into the cli binary.

## Needs

* Create an account on Maxmind site for downloading database.

## Build

* Makefile download and extract maxmind geolite database.
* Makefile get exact asset name and replace in template init.go
* ~~Makefile build name.go from assets with go-bindata.~~
* Asset is embedded with go embed facility

## Usage

* "-V" is used to show maxmind geolite database version.
* "-r" option read stdin and try to resolve dns in parallel.
* otherwise arg1 is checked.

```Shell
 $/local/bin/geocli -V 
/local/bin/geocli build with GeoLite2-Country_20191001/GeoLite2-Country.mmdb

 # rg "imap\[.* login: " /var/log/imapd.log | cut -d: -f5- | sed "s/ User.*//g" | rg -v "webmail|TLS" | sort -u|rg -v uvsq | cut -d\[ -f2 | cut -d\] -f1 | sor
37.171.85.185 [37-171-85-185.coucou-networks.fr.] France, FR
10.172.16.103 [unknown] , 
37.164.163.41 [unknown] France, FR
37.170.73.3 [37-170-73-3.coucou-networks.fr.] France, FR
37.164.246.169 [unknown] France, FR
176.179.79.3 [176-179-79-3.abo.bbox.fr.] France, FR
37.170.3.114 [37-170-3-114.coucou-networks.fr.] France, FR
37.171.67.201 [37-171-67-201.coucou-networks.fr.] France, FR
```

## References

* go-bindata > go get -u github.com/jteeuwen/go-bindata/...

* go embed see [golang embed](https://golang.org/pkg/embed/)
