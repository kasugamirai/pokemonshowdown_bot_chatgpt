package commands_test

import (
	"strings"
	"testing"

	"xy.com/pokemonshowdownbot/commands"
	"xy.com/pokemonshowdownbot/config"
	"xy.com/pokemonshowdownbot/database"
)

func TestPrompt(t *testing.T) {
	config.LoadConfig("../config.json")

	msg := "What is the capital of France?"

	response, err := commands.Prompt(msg)
	if err != nil {
		t.Fatalf("Error in Prompt: %v", err)
	}

	expectedResponse := "Paris"
	if response != expectedResponse {
		t.Errorf("Expected response '%s', but got '%s'", expectedResponse, response)
	}
}

func TestAddSticker(t *testing.T) {
	config.LoadConfig("../config.json")
	err := database.InitDB()
	if err != nil {
		t.Fatalf("Failed to initialize database: %v", err)
	}

	// Replace with your own example data for testing
	name := "example_sticker_name"
	url := "https://example.com/sticker.jpg"

	msg := name + " " + url
	response, err := commands.AddSticker(msg)
	if err != nil {
		t.Fatalf("Error in AddSticker: %v", err)
	}

	expectedResponse := "The image has been successfully added to the database."
	if response != expectedResponse {
		t.Errorf("Expected response '%s', but got '%s'", expectedResponse, response)
	}
}

func TestFindStickerByName(t *testing.T) {
	config.LoadConfig("../config.json")
	err := database.InitDB()
	if err != nil {
		t.Fatalf("Failed to initialize database: %v", err)
	}

	// Replace with the name of a sticker that has been added to the database
	name := "example_sticker_name"

	response, err := commands.FindStickerByName(name)
	if err != nil {
		t.Fatalf("Error in FindStickerByName: %v", err)
	}

	expectedSubstring := "<img src=\"https://example.com/sticker.jpg\" height=\"100\">"
	if !strings.Contains(response, expectedSubstring) {
		t.Errorf("Expected response to contain '%s', but got '%s'", expectedSubstring, response)
	}
}
