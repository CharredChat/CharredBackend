package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/CharredChat/CharredBackend/routes"

	"github.com/CharredChat/charrid/CharredProcess"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	var proc = CharredProcess.New(127)
	fmt.Println("Server starting.")
	var test, err = proc.Generate()
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("ID: %X\n", test.AsUint64())
	setupRoutes()
	http.ListenAndServe(":8080", nil)
}

func setupRoutes() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/ws", routes.WsEndpoint)
}
