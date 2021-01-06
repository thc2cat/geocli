package main

import "testing"

type teststruct struct {
	name string
	have string
	want string
}

var tests = []teststruct{
	{"soleil", "193.51.24.1", "193.51.24.1 [soleil.uvsq.fr.] France, FR"},
	{"network", "193.51.24.0", "193.51.24.0 [unknown] France, FR"},
	{"null", "0.0.0.0", "0.0.0.0 [unknown] , "},
	{"litterals", "should not work", ""}, // produce log errors
}

func Test_parse(t *testing.T) {
	db := initdb()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseandprint(tt.have, db); got != tt.want {
				t.Errorf("parseandprint() = |%v|, want |%v|", got, tt.want)
			}
		})
	}

	db.Close()
}
