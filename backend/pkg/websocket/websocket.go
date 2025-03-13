package websocket

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// Upgrader is used to upgrade the HTTP connection to a WebSocket connection
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,

	// CheckOrigin is used to check the origin of the WebSocket connection
	// We are allowing all origins
	CheckOrigin: func(r *http.Request) bool { return true },
}

func Upgrade(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return conn, nil
}

func Reader(conn *websocket.Conn) {
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		fmt.Println(string(p))

		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}
	}
}

func Writer(conn *websocket.Conn) {
	for {
		fmt.Println("Sending")
		messageType, message, err := conn.NextReader()
		if err != nil {
			log.Println(err)
			return
		}

		w, err := conn.NextWriter(messageType)
		if err != nil {
			log.Println(err)
			return
		}

		if _, err := io.Copy(w, message); err != nil {
			log.Println(err)
			return
		}

		if err := w.Close(); err != nil {
			fmt.Println(err)
			return
		}
	}
}
