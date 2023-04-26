package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/CharredChat/charrid/CharredProcess"
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
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
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		conn, _, _, err := ws.UpgradeHTTP(r, w)
		if err != nil {
			log.Fatal(err)
			return
		}
		go func() {
			defer conn.Close()

			for {
				msg, op, err := wsutil.ReadClientData(conn)
				if err != nil {
					// handle error
					log.Fatal(err)
					return
				}
				err = wsutil.WriteServerMessage(conn, op, msg)
				if err != nil {
					// handle error
					log.Fatal(err)
					return
				}
			}
		}()
	})
	http.ListenAndServe("127.0.0.1:8080", nil)

}
