package routes

import (
	"log"
	"net/http"

	"github.com/CharredChat/CharredBackend/comm"
	"github.com/CharredChat/charrid/CharredProcess"
	"github.com/gorilla/websocket"
)

var Proc CharredProcess.CharredProcess

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
	var incoming comm.Message
	for {
		err := conn.ReadJSON(&incoming)
		if err != nil {
			log.Println(err)
			return
		}

		test, err := Proc.Generate()
		if err != nil {
			log.Println(err)
			return
		}

		log.Println(test.AsUint64())

		switch incoming.Op {
		case 1:
			sayMsg, ok := incoming.OpData.(map[string]interface{})
			if !ok {
				return
			}
			sayStr, ok := sayMsg["say"].(string)
			if !ok {
				return
			}
			say := comm.OpOne{Say: sayStr}

			log.Println(say.Say)
		}

		log.Println(incoming.Op)
		err = conn.WriteJSON(incoming)
		if err != nil {
			log.Println(err)
			return
		}
	}
}
