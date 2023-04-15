package chatgpt_test

import (
	"strings"
	"testing"

	"xy.com/pokemonshowdownbot/chatgpt"
	"xy.com/pokemonshowdownbot/config"
)

func TestChatWithGPT(t *testing.T) {
	config.LoadConfig("../config.json")
	prompt := "What is the capital of France?"

	response, err := chatgpt.ChatWithGPT(prompt)
	if err != nil {
		t.Fatalf("Error in ChatWithGPT: %v", err)
	}

	expectedResponse := "Paris"
	if !strings.Contains(response, expectedResponse) {
		t.Errorf("Expected response to contain '%s', but got '%s'", expectedResponse, response)
	}
}
