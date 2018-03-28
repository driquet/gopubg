package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/driquet/gopubg"
	"github.com/sirupsen/logrus"
)

var (
	input string
)

func usage() {
	fmt.Printf("usage: %s [options]\n", os.Args[0])
	flag.PrintDefaults()
	os.Exit(1)
}

func init() {
	// Parameters
	flag.StringVar(&input, "input", "", "input file")

	// Parse parameters
	flag.Parse()

	// Verify parameters
	if input == "" {
		usage()
	}
}

func main() {
	file, err := os.Open(input)
	if err != nil {
		logrus.Fatal(err)
	}

	t, err := gopubg.ParseTelemetry(file)
	if err != nil {
		logrus.Fatal(err)
	}

	fmt.Printf("%d events parsed\n", len(t.Events))
	fmt.Printf("%d players\n", len(t.PlayerNames))

	for _, p := range t.PlayerNames {
		fmt.Printf(" - %s\n", p)
	}
}
