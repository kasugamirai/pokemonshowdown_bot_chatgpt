package config

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

var (
	// Instance of Config struct, accessible through the package
	Instance Config
)

type Config struct {
	DatabaseDSN     string
	Chatgpt         chatgpt
	Pokemonshowdown pokemonshowdown
}

type pokemonshowdown struct {
	Server   string
	Username string
	Password string
	Avatar   string
	Room     string
}

type chatgpt struct {
	OpenAIAPIURL   string
	Model          string
	OPENAI_API_KEY string
}

func LoadConfig(Path string) {
	// Load the configuration file
	configFile, err := os.Open(Path)
	if err != nil {
		log.Printf("Unable to open config file: %v, using defaults.", err)
		return
	}
	defer configFile.Close()

	// Read the configuration file
	bytes, err := io.ReadAll(configFile)
	if err != nil {
		log.Printf("Unable to read config file: %v, using defaults.", err)
		return
	}

	// Unmarshal the JSON configuration into Config struct
	err = json.Unmarshal(bytes, &Instance)
	if err != nil {
		log.Printf("Unable to parse config file: %v, using defaults.", err)
		return
	}
}
