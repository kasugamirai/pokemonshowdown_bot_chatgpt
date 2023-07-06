package config

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

var (
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
	configFile, err := os.Open(Path)
	if err != nil {
		log.Printf("Unable to open config file: %v, using defaults.", err)
		return
	}
	defer configFile.Close()

	bytes, err := io.ReadAll(configFile)
	if err != nil {
		log.Printf("Unable to read config file: %v, using defaults.", err)
		return
	}

	err = json.Unmarshal(bytes, &Instance)
	if err != nil {
		log.Printf("Unable to parse config file: %v, using defaults.", err)
		return
	}

	// get pokemonshowdown password and OpenAI API key from environment variables
	Instance.Pokemonshowdown.Password = os.Getenv("POKEMONSHOWDOWN_PASSWORD")
	Instance.Chatgpt.OPENAI_API_KEY = os.Getenv("OPENAI_API_KEY")

	if Instance.Pokemonshowdown.Password == "" || Instance.Chatgpt.OPENAI_API_KEY == "" {
		log.Println("Some environment variables are not set. Make sure to set POKEMONSHOWDOWN_PASSWORD and OPENAI_API_KEY.")
	}
}
