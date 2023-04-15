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

	conn, err := showdown.ConnectToServer(ps.Server)
	if err != nil {
		log.Fatal("Error connecting to the server:", err)
	}
	defer conn.Close()

	showdown.Login(conn, ps.Username, ps.Password, ps.Avatar)
	showdown.JoinRoom(conn, ps.Room)
	go showdown.ReadMessages(conn)

	for {
		time.Sleep(1 * time.Minute)
	}
}
