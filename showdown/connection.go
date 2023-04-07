package showdown

import (
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/gorilla/websocket"
)

var dialer = websocket.Dialer{
	Proxy: http.ProxyFromEnvironment,
}

func ConnectToServer(server string) (*websocket.Conn, error) {
	u := url.URL{Scheme: "ws", Host: server, Path: "/showdown/websocket"}
	conn, _, err := dialer.Dial(u.String(), nil)
	return conn, err
}

func ReadMessages(conn *websocket.Conn) {
	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error reading message:", err)
			return
		}
		fmt.Printf("%d: %s\n", messageType, message)
	}
}
