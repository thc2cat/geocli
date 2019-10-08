package main

// Historique:
// 1.1 - limitation a 36 requetes DNS en //
// 1.2 - trim space before resolving dns
// 1.3 - less fatal errors
//
import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"sync"

	geoip2 "github.com/oschwald/geoip2-golang"
)

var maxrequests = 64

func main() {

	if len(os.Args) < 2 { // Need an ip
		fmt.Printf("Usage: %s ip\n", os.Args[0])
		os.Exit(-1)
	}
	if os.Args[1] == "-v" {
		fmt.Printf("%s is using %s\n", os.Args[0], Assetname)
		os.Exit(0)
	}

	db := initdb()
	defer db.Close()

	// If you are using strings that may be invalid, check that ip is not nil
	//ip := net.ParseIP("193.49.159.7")

	if os.Args[1] != "-r" {
		fmt.Printf("%s\n", parseandprint(os.Args[1], db))
		os.Exit(0)
	}
	// on va lire l'entree standard, faire les requetes DNS en //
	// puis afficher chaque entree
	readandprintbulk(db)
	os.Exit(0)

}

func initdb() *geoip2.Reader {
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
	return db
}

func parseandprint(ips string, db *geoip2.Reader) string {
	ip := net.ParseIP(ips)
	if ip == nil {
		log.Printf("Unable to parse ip : \"%s\"", ips)
		return ""
	}
	record, err := db.Country(ip)
	if err != nil {
		log.Printf("Unable to geoloc %s", ips)
		return ""
	}

	output := ips

	addrs, err := net.LookupAddr(ips)
	if err == nil {
		// log.Fatal("Resolution error", err.Error())
		for _, s := range addrs {
			// fmt.Print(" [", s)
			output += " [" + s
		}
		// fmt.Printf("]")
		output += "] "
	} else {
		// fmt.Print(" [unknown]")
		output += " [unknown] "
	}
	// fmt.Printf(" %s, %s\n", record.Country.Names["en"], record.Country.IsoCode)
	output += record.Country.Names["en"] + ", " + record.Country.IsoCode
	return output
}

func readandprintbulk(db *geoip2.Reader) {
	var line string
	// var bulk []string
	var wg sync.WaitGroup

	var limitChan = make(chan bool, maxrequests)

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {

		line = scanner.Text()
		line = strings.TrimSpace(line)
		// bulk = append(bulk, line)
		wg.Add(1)
		limitChan <- true
		go func(line string, mywg *sync.WaitGroup, mychan chan bool) {
			// parseandprint(line, db)
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
