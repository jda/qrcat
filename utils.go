package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/rsc/qr"
)

func getInputData() (f *os.File, err error) {
	switch flag.NArg() {
	case 0:
		f = os.Stdin
	case 1:
		f, err = os.Open(flag.Arg(0))
	default:
		err = fmt.Errorf("missing file input")
	}

	return f, err
}

func getRedundancyLevel(r string) (q qr.Level, err error) {
	r = strings.ToUpper(r)

	switch r {
	case "L":
		q = qr.L
	case "M":
		q = qr.M
	case "Q":
		q = qr.Q
	case "H":
		q = qr.H
	default:
		err = fmt.Errorf("%s it not a valid QR code redundancy level", r)
	}

	return q, err
}
