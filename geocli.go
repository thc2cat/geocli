package main

import (
	"fmt"
	"log"
	"net"
	"os"

	geoip2 "github.com/oschwald/geoip2-golang"
)

func main() {

	data, err := Asset(Assetname)
	if err != nil {
		// Asset was not found.
		log.Fatal("Error opening Asset ", Assetname)
	}

	//Instead of opening the file, embedding it with
	//Alternative db, err := geoip2.Open("/local/etc/GeoLite2-Country.mmdb")
	db, err := geoip2.FromBytes(data)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	// If you are using strings that may be invalid, check that ip is not nil
	//ip := net.ParseIP("193.49.159.7")
	ip := net.ParseIP(os.Args[1])
	if ip == nil {
		log.Fatal("Unable to parse ip ", os.Args[1])
	}
	record, err := db.Country(ip)
	if err != nil {
		log.Fatal("Unable to geoloc ", os.Args[1])
	}

	fmt.Print(os.Args[1])

	addrs, err := net.LookupAddr(os.Args[1])
	if err == nil {
		// log.Fatal("Resolution error", err.Error())
		for _, s := range addrs {
			fmt.Print(" [", s)
		}
		fmt.Printf("]")
	} else {
		fmt.Print(" [unknown]")
	}

	fmt.Printf(" %s, %s\n", record.Country.Names["en"], record.Country.IsoCode)

}
