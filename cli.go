package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"io"
	"os"
	"time"

	qr "github.com/mdp/qrterminal"
)

func main() {
	flag.Usage = usage
	displayInterval := flag.Duration("d", 5*time.Second, "duration to show each code")
	redundancy := flag.String("r", "L", "redundancy level [L|M|Q|H]")
	blockSize := flag.Int("s", 2048, "size of code chunk in bytes")
	halfBlock := flag.Bool("half", false, "output using unicode half blocks")
	flag.Parse()

	qrRedundancy, err := getRedundancyLevel(*redundancy)
	if err != nil {
		usage()
	}

	in, err := getInputData()
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not process input: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("display interval would be %d\n", *displayInterval)

	block := 1
	b := make([]byte, *blockSize)
	for {
		fmt.Printf("block %d\n", block)
		block++

		n, err := in.Read(b)
		if err != nil {
			if err != io.EOF {
				fmt.Fprintf(os.Stderr, "could not read input: %s\n", err)
				os.Exit(1)
			}
		}

		out := base64.StdEncoding.EncodeToString(b)
		if *halfBlock == true {
			qr.GenerateHalfBlock(out, qrRedundancy, os.Stdout)
		} else {
			qr.Generate(out, qrRedundancy, os.Stdout)
		}

		if n < *blockSize { // read less than size so we're done
			break
		}
		time.Sleep(*displayInterval)
	}

}

func usage() {
	fmt.Printf("Usage: %s [OPTIONS] FILENAME\n", os.Args[0])
	flag.PrintDefaults()
	os.Exit(1)
}
