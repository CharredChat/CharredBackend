package main

import (
	"fmt"
	"log"

	"github.com/CharredChat/charrid/CharredProcess"
)

func main() {
	var proc = CharredProcess.New(127)
	fmt.Println("Server starting.")
	var test, err = proc.Generate()
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("ID: %X\n", test.AsUint64())
}