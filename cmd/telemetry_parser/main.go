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
	logrus.SetLevel(logrus.DebugLevel)

	file, err := os.Open(input)
	if err != nil {
		logrus.Fatal(err)
	}

	t, err := gopubg.ParseTelemetry(file)
	if err != nil {
		logrus.Fatal(err)
	}

	fmt.Printf("%d events parsed\n", len(t.Events))
	fmt.Printf("%d players\n", len(t.Players))

	for _, player := range t.Players {
		fmt.Printf(" - %s (%d events, ranking=%d)\n", player.Name, len(player.Events), player.Ranking)
		if player.Ranking == 1 {
			fmt.Printf("winner: %s\n", player.Name)

			for idx, evt := range player.Events {
				fmt.Printf("%d: %s\n", idx, gopubg.KnownEventTypes[evt.Type])
			}
			break
		}
	}
}
