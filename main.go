package main

import (
	"log"
	"time"

	"xy.com/pokemonshowdownbot/commands"
	"xy.com/pokemonshowdownbot/config"
	"xy.com/pokemonshowdownbot/database"
	"xy.com/pokemonshowdownbot/showdown"
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

	// Infinite loop to keep the program running
	for {
		// Connect to server
		conn, err := showdown.ConnectToServer(ps.Server)
		if err != nil {
			log.Println("Error connecting to the server:", err)
			time.Sleep(5 * time.Second) // Wait for 5 seconds before retrying
			continue
		}

		showdown.Login(conn, ps.Username, ps.Password, ps.Avatar)
		showdown.JoinRoom(conn, ps.Room)

		// Loop to read messages
		for {
			_, _, err := conn.ReadMessage()
			if err != nil {
				log.Println("Error reading message, lost connection to the server:", err)
				conn.Close() // Close the connection before retrying
				break
			}

			// Process the message...
		}

		time.Sleep(5 * time.Second) // Wait for 5 seconds before retrying
	}
}
