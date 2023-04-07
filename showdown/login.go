package showdown

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/gorilla/websocket"
)

func Login(conn *websocket.Conn, username, password, avatar string) {
	challstr, err := GetChallstr(conn)
	if err != nil {
		log.Println("Error getting challstr:", err)
		return
	}

	assertion, err := getAssertion(username, password, challstr)
	if err != nil {
		log.Println("Error getting assertion:", err)
		return
	}

	conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("|/trn %s,0,%s\n", username, assertion)))
	time.Sleep(1 * time.Second)
	conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("|/avatar %s\n", avatar)))
}

func GetChallstr(conn *websocket.Conn) (string, error) {
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			return "", err
		}

		parts := strings.Split(string(message), "|")
		if len(parts) > 2 && parts[1] == "challstr" {
			return parts[2] + "|" + parts[3], nil
		}
	}
}

func getAssertion(username, password, challstr string) (string, error) {
	data := url.Values{}
	if password != "" {
		data.Set("act", "login")
		data.Set("name", username)
		data.Set("pass", password)
	} else {
		data.Set("act", "getassertion")
		data.Set("userid", username)
	}
	data.Set("challstr", challstr)

	resp, err := http.PostForm("https://play.pokemonshowdown.com/action.php", data)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if password != "" {
		var jsonResponse map[string]interface{}
		err = json.Unmarshal(body[1:], &jsonResponse)
		if err != nil {
			return "", err
		}

		if jsonResponse["actionsuccess"].(bool) {
			return jsonResponse["assertion"].(string), nil
		} else {
			return "", fmt.Errorf("login unsuccessful")
		}
	} else {
		return string(body), nil
	}
}
