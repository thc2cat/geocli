package main

import (
	_ "embed"
	"log"

	geoip2 "github.com/oschwald/geoip2-golang"
)

//go:embed %%GEOIPDATANAME%%
var embeddeddata []byte

func initdb() *geoip2.Reader {

	// 	data, err := Asset(Assetname)
	// 	if err != nil {
	// 		// Asset was not found.
	// 		log.Fatal("Error opening Asset ", Assetname)
	// 	}
	// 	// Memo : Instead of opening the file, embedding it with
	// 	// db, err := geoip2.Open("/local/etc/GeoLite2-Country.mmdb")

	db, err := geoip2.FromBytes(embeddeddata)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

var Assetname = "%%GEOIPDATANAME%%"
