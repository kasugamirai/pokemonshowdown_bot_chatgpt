# pokemonshowdown_bot_chatgpt

Pokémon Showdown Bot with ChatGPT Integration

This project is a Pokémon Showdown chatbot capable of connecting to a game server, logging in, joining specified chat rooms, and responding to commands. The chatbot is written in Go and uses the Gorilla WebSocket library for real-time communication with the server. Additionally, the bot is integrated with OpenAI's ChatGPT, providing enhanced natural language understanding and response generation capabilities.

Getting Started

Prerequisites
Go programming language installed
Gorilla WebSocket library
OpenAI API Key
Installation

Clone the repository:
git clone https://github.com/kasugamirai/pokemonshowdown_bot_chatgpt.git

Change to the project directory:
cd pokemonshowdownbot

Install the Gorilla WebSocket library:
go get -u github.com/gorilla/websocket

Usage
Edit the config.json file to configure the server, username, password, avatar, room, and OpenAI API key.

Build the project:
go build

Run the compiled binary:
./pokemonshowdownbot

Features
Connects to the Pokémon Showdown server
Logs in with a given username, password, and avatar
Joins specified chat rooms
Reads and responds to messages in real-time
Integrates with OpenAI's ChatGPT for advanced natural language understanding and response generation
Code Structure
main.go: Main entry point for the application
showdown/: Contains the core functionality for connecting to the Pokémon Showdown server and interacting with chat rooms
chatgpt/: Contains the code for interacting with the OpenAI API and handling ChatGPT integration

