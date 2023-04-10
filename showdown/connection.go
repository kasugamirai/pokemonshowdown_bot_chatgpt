package showdown

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/gorilla/websocket"
	"xy.com/pokemonshowdownbot/chatgpt"
)

var dialer = websocket.Dialer{
	Proxy: http.ProxyFromEnvironment,
}

func ConnectToServer(server string) (*websocket.Conn, error) {
	u := url.URL{Scheme: "ws", Host: server, Path: "/showdown/websocket"}
	conn, _, err := dialer.Dial(u.String(), nil)
	return conn, err
}

func ReadMessages(conn *websocket.Conn, room string) {
	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error reading message:", err)
			return
		}
		fmt.Printf("%d: %s\n", messageType, message)
		msg := string(message)
		parts := strings.Split(msg, "|")
		if len(parts) > 4 && IsStaff(parts[3]) && strings.HasPrefix(parts[4], ".prompt") {
			response, err := chatgpt.ChatWithGPT(parts[4][7:])
			if err != nil {
				log.Println("Error ChatWithGPT message:", err)
				return
			}
			if len(response) < 218 {
				SendMessage(conn, room, response)
			} else {
				SendMessage(conn, room, "!code "+response)
			}

		}
	}
}

func IsStaff(username string) bool {
	auth := [3]string{"@", "#", "~"}
	for _, a := range auth {
		if strings.HasPrefix(username, a) {
			return true
		}

	}
	return false
}

func SendMessage(conn *websocket.Conn, room, message string) {
	message = strings.ReplaceAll(message, "\n", "")
	conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("%s|%s\n", room, message)))
}
