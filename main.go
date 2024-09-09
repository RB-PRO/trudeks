package main

import (
	"fmt"
	tgz4b "trudeks/internal/tg_z4b"
)

func main() {
	defer recoverPanic()

	//tgz4b.Start()
	tgz4b.StartLocal()

	fmt.Println("gotovo blyat")
	fmt.Scanf(" ")

	//tgz4b.ParsingCounterGoRoutines()
	//parsing.MSK()
	//parsing.MO()
	//parsing.MO2()
	//tgapp, err := app.New()
	//if err != nil {
	//	panic(err)
	//}
	//log.Fatal(tgapp.Run())
}

func recoverPanic() {
	if err := recover(); err != nil {
		fmt.Printf("RECOVERED: %v\n", err)
		fmt.Scanf(" ")
	}
}
