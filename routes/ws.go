package routes

import (
	"log"
	"net/http"

	"github.com/CharredChat/CharredBackend/transactions"
	"github.com/CharredChat/charrid/CharredProcess"
	"github.com/gorilla/websocket"
	"google.golang.org/protobuf/proto"
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
	for {
		incoming := &transactions.Request{}
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		err = proto.Unmarshal(message, incoming)
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

		switch op := incoming.Op.(type) {
		case *transactions.Request_AuthRequest_:
			sayStr := op.AuthRequest.GetSay()
			log.Println(sayStr)

			var resp = transactions.Response{
				Op: &transactions.Response_MessageCreatedResponse_{
					MessageCreatedResponse: &transactions.Response_MessageCreatedResponse{
						Id: test.AsUint64(),
					},
				},
			}

			respdata, err := proto.Marshal(&resp)
			if err != nil {
				log.Println(err)
				return
			}

			err = conn.WriteMessage(websocket.BinaryMessage, respdata)
			if err != nil {
				log.Println(err)
				return
			}
		}
	}
}
