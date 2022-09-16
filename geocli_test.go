package main

import (
	"io/ioutil"
	"log"
	"os"
	"testing"
)

type teststruct struct {
	name string
	have string
	want string
}

var tests = []teststruct{
	{"soleil", "193.51.24.1", "193.51.24.1 [soleil.uvsq.fr.] France, FR, local"},
	{"network", "193.51.24.0", "193.51.24.0 [unknown] France, FR, local"},
	{"Private ip", "192.168.0.0", "192.168.0.0 [unknown] , local"},
	{"not local", "101.68.211.3", "101.68.211.3 [unknown] China, CN"},
	{"litterals", "should not work", ""}, // produce log errors
}

func Test_parse(t *testing.T) {
	db := initdb()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseandprint(tt.have, db); got != tt.want {
				t.Errorf("parseandprint(%s) => |%v| !=  expected |%v|", tt.have, got, tt.want)
			}
		})
	}

	db.Close()
}

func Test_readandprintbulk(t *testing.T) {
	content := []byte("193.51.24.1")
	tmpfile, err := ioutil.TempFile("", "testfile_tmp")
	if err != nil {
		log.Fatal(err)
	}

	db := initdb()

	defer os.Remove(tmpfile.Name()) // clean up

	if _, err := tmpfile.Write(content); err != nil {
		log.Fatal(err)
	}

	if _, err := tmpfile.Seek(0, 0); err != nil {
		log.Fatal(err)
	}

	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }() // Restore original Stdin

	os.Stdin = tmpfile
	readandprintbulk(db)

	if err := tmpfile.Close(); err != nil {
		log.Fatal(err)
	}

}
