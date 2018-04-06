package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/driquet/gopubg"
	"github.com/sirupsen/logrus"
)

var (
	key        string
	playerName string
	shard      string
)

func usage() {
	fmt.Printf("usage: %s [options]\n", os.Args[0])
	flag.PrintDefaults()
	os.Exit(1)
}

func init() {
	// Parameters
	flag.StringVar(&key, "key", "", "api key")
	flag.StringVar(&playerName, "name", "", "player name")
	flag.StringVar(&shard, "shard", "", "shard")

	// Parse parameters
	flag.Parse()

	// Verify parameters
	if key == "" || playerName == "" || shard == "" {
		usage()
	}
}

func main() {
	api := gopubg.NewAPI(key)
	_, err := api.RequestSinglePlayerByName(shard, playerName)
	if err != nil {
		logrus.Fatal(err)
	}
}
