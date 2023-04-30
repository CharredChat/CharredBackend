package routes

import (
	"log"
	"net/http"

	"github.com/CharredChat/CharredBackend/comm"
	"github.com/gorilla/websocket"
)

// trunk-ignore(gitleaks/generic-api-key)
const MASTER_DEMO_TOKEN = "MTEwMjgwMjU0MjY4NDA3Nzc4.MTY4Mjc0MzEzNzY0MA==.c4t51kFWSEmdmaPnKoyUuu8E78E"

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func WsEndpoint(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}
	log.Println("Client successfully connected.")
	reader(ws)
	ws.Close()
}

func reader(conn *websocket.Conn) {
	var json comm.Message
	for {
		err := conn.ReadJSON(&json)
		if err != nil {
			log.Println(err)
			return
		}

		log.Println(json.Op)
		err = conn.WriteJSON(json)
		if err != nil {
			log.Println(err)
			return
		}
	}
}
