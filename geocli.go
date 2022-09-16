package main

// Revisions :
// 1.0 - embedding geolite2Country data in binary code
// 1.1 - limit // DNS requests
// 1.2 - trim space before resolving dns
// 1.3 - less fatal errors
// 1.4 - cleaner main() code and parse tests
// 1.5 - using go embed

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"regexp"
	"strings"
	"sync"

	geoip2 "github.com/oschwald/geoip2-golang"
)

var (
	maxrequests = 512
	// Version given by git tag via Makefile
	Version  string
	Privates = regexp.MustCompile(`^(10|172\.(1[6789]|2[0-9]|3[01])|192\.168|193.51\.(2[456789]|3[0-9]|4[12]))\.`)
)

func main() {

	switch {
	case len(os.Args) < 2:
		fmt.Printf("Usage: %s ip|-r(ead stdin)|-V(ersion)\n", os.Args[0])
		os.Exit(-1)
	case os.Args[1] == "-V":
		fmt.Printf("%s build with %s\n", Version, Assetname)
		os.Exit(0)
	}

	db := initdb()
	defer db.Close()

	// If you are using strings that may be invalid, check
	// that ip is not nil
	//ip := net.ParseIP("193.49.159.7")

	switch {
	case os.Args[1] == "-r":
		readandprintbulk(db)
	default:
		fmt.Printf("%s\n", parseandprint(os.Args[1], db))
	}

}

func parseandprint(ips string, db *geoip2.Reader) string {
	var record *geoip2.Country

	ip := net.ParseIP(ips)
	if ip == nil {
		log.Printf("Unable to parse ip : \"%s\"", ips)
		return ""
	}

	var err error
	record, err = db.Country(ip)
	if err != nil {
		log.Printf("Unable to geoloc \"%s\"", ips)
		return ""
	}

	output := ips

	addrs, err := net.LookupAddr(ips)
	if err == nil {
		for _, s := range addrs {
			output += " [" + s
		}
		output += "] "
	} else {
		output += " [unknown] "
	}

	if record != nil && record.Country.Names["en"] != "" {
		output += record.Country.Names["en"] + ", " + record.Country.IsoCode
	}
	if Privates.MatchString(ips) {
		output += ", local"
	}
	return output
}

func readandprintbulk(db *geoip2.Reader) {
	var line string
	var wg sync.WaitGroup
	var limitChan = make(chan bool, maxrequests)

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line = strings.TrimSpace(scanner.Text())

		wg.Add(1)
		limitChan <- true // will block after maxrequests

		go func(line string, mywg *sync.WaitGroup, mychan chan bool) {
			out := parseandprint(line, db)
			if len(out) > 0 {
				fmt.Printf("%s\n", out)
			}
			<-mychan
			mywg.Done()
		}(line, &wg, limitChan)

	}
	wg.Wait()
}
