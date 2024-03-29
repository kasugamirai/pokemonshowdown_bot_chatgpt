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

	// Auto reconnect loop
	for {
		conn, err := showdown.ConnectToServer(ps.Server)
		if err != nil {
			log.Println("Error connecting to the server:", err)
			time.Sleep(5 * time.Second) // Wait for 5 seconds before retry
			continue
		}

		// If we reach here, we are connected
		defer conn.Close()

		showdown.Login(conn, ps.Username, ps.Password, ps.Avatar)
		showdown.JoinRoom(conn, ps.Room)

		var done = make(chan bool)

		// Read messages in a separate goroutine to avoid blocking
		go func() {
			showdown.ReadMessages(conn)
			done <- true
		}()

		// Wait until the ReadMessages function finishes running
		<-done
	}
}
