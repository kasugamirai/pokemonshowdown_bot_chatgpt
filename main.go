package main

import (
	"log"
	"time"

	"xy.com/pokemonshowdownbot/commands"
	"xy.com/pokemonshowdownbot/config"
	"xy.com/pokemonshowdownbot/database"
	"xy.com/pokemonshowdownbot/showdown"

	"github.com/gorilla/websocket"
)

func main() {
	// Load configuration
	config.LoadConfig("config.json")
	var ps = config.Instance.Pokemonshowdown
	// Initialize the database connection
	err := database.InitDB()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	// Load commands
	commands.LoadCommands()

	// Infinite loop to handle reconnecting
	for {
		conn, err := showdown.ConnectToServer(ps.Server)
		if err != nil {
			log.Println("Error connecting to the server, retrying in 1 minute:", err)
			time.Sleep(1 * time.Minute)
			continue
		}

		// Set the Ping handler
		conn.SetPingHandler(func(string) error {
			conn.WriteMessage(websocket.PongMessage, []byte{})
			return nil
		})

		showdown.Login(conn, ps.Username, ps.Password, ps.Avatar)
		showdown.JoinRoom(conn, ps.Room)

		// Start a goroutine to read messages
		go showdown.ReadMessages(conn)

		// Start a goroutine to send Ping messages periodically
		go func() {
			for {
				time.Sleep(5 * time.Minute)
				if err := conn.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
					log.Println("Error sending ping, reconnecting:", err)
					conn.Close()
					return
				}
			}
		}()

		// Wait here until the connection is closed
		for {
			_, _, err := conn.ReadMessage()
			if err != nil {
				log.Println("Error reading message, reconnecting:", err)
				conn.Close()
				break
			}
		}

		log.Println("Reconnecting in 1 minute...")
		time.Sleep(1 * time.Minute)
	}
}
