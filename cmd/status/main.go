package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/driquet/gopubg"
	"github.com/sirupsen/logrus"
)

var (
	key string
)

func usage() {
	fmt.Printf("usage: %s [options]\n", os.Args[0])
	flag.PrintDefaults()
	os.Exit(1)
}

func init() {
	// Parameters
	flag.StringVar(&key, "key", "", "api key")

	// Parse parameters
	flag.Parse()

	// Verify parameters
	if key == "" {
		usage()
	}
}

func main() {
	api := gopubg.NewAPI(key)
	err := api.RequestStatus()
	if err != nil {
		logrus.Fatal(err)
	}
}
