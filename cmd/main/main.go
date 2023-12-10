package main

import (
	"log"

	app "github.com/RB-PRO/trudeks/internal/tgcouterapp"
)

func main() {
	// tgz4b.Start()
	// tgz4b.ParsingCounterGoRoutines()
	// parsing.MSK()
	// parsing.MO()
	// parsing.MO2()
	tgapp, err := app.New()
	if err != nil {
		panic(err)
	}
	log.Fatal(tgapp.Run())
}
