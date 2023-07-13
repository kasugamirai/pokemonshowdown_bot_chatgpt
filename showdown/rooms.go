package showdown

import (
	"fmt"
	"log"
	"strings"

	"github.com/gorilla/websocket"
	"xy.com/pokemonshowdownbot/commands"
)

var auth = map[string]struct{}{
	"+": struct{}{},
	"%": struct{}{},
	"@": struct{}{},
	"#": struct{}{},
	"~": struct{}{},
}

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
			processCommand(conn, parts)
		}
	}
}

func processCommand(conn *websocket.Conn, parts []string) {
	for s, fn := range commands.CommandsMap {
		if strings.HasPrefix(parts[4], "."+s) {
			res, err := fn(parts[4][len(s)+2:])
			if err != nil {
				log.Println(err)
				continue
			}
			SendMessage(conn, parts[0][1:len(parts[0])-1], res)
			break
		}
	}
}

func IsStaff(username string) bool {
	_, found := auth[string(username[0])]
	return found
}

func SendMessage(conn *websocket.Conn, room, message string) {
	message = strings.ReplaceAll(message, "\n", "")
	conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("%s|%s\n", room, message)))
}
