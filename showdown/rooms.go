package showdown

import (
	"fmt"
	"strings"

	"github.com/gorilla/websocket"
)

func JoinRoom(conn *websocket.Conn, room string) {
	conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("|/join %s\n", room)))
}

func SendMessage(conn *websocket.Conn, room, message string) {
	message = strings.ReplaceAll(message, "\n", "")
	conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("%s|%s\n", room, message)))
}
