package showdown

import (
	"fmt"
	"log"
	"strings"

	"github.com/gorilla/websocket"
	"xy.com/pokemonshowdownbot/commands"
)

func JoinRoom(conn *websocket.Conn, room string) {
	conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("|/join %s\n", room)))
}

func ReadMessages(conn *websocket.Conn) {
	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error reading message:", err)
			return
		}
		fmt.Printf("%d: %s\n", messageType, message)
		msg := string(message)
		parts := strings.Split(msg, "|")
		if len(parts) > 4 && parts[4][0] == '.' && IsStaff(parts[3]) {
			for s, fn := range commands.CommandsMap {
				if strings.HasPrefix(parts[4], "."+s) {
					res, err := fn(parts[4][len(s)+2:])
					if err != nil {
						log.Println(err)
						break
					}
					SendMessage(conn, parts[0][1:len(parts[0])-1], res)
					break
				}
			}
		}
	}
}

func IsStaff(username string) bool {
	auth := [5]string{"+", "%", "@", "#", "~"}
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
