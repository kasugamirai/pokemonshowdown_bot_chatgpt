package main

import (
	"log"
	"time"

	"xy.com/pokemonshowdownbot/showdown"
)

const (
	server   = "sim.smogon.com:8000"
	username = "-"
	password = "-"
	avatar   = "dawn"
	room     = "chinese"
)

func main() {
	conn, err := showdown.ConnectToServer(server)
	if err != nil {
		log.Fatal("Error connecting to the server:", err)
	}
	defer conn.Close()

	showdown.Login(conn, username, password, avatar)
	showdown.JoinRoom(conn, room)
	go showdown.ReadMessages(conn, room)

	for {
		time.Sleep(1 * time.Minute)
	}
}
