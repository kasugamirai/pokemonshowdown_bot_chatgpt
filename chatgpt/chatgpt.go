// Package chatgpt provides an interface to interact with OpenAI's GPT-3.5-turbo model.
package chatgpt

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"xy.com/pokemonshowdownbot/config"
)

// Response represents the structure of the response from the OpenAI API.
type Response struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}

// Message is a structure representing a message sent to the API.
type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// ChatCompletionRequest is the structure of the request sent to the OpenAI API.
type ChatCompletionRequest struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

// ChatWithGPT sends a prompt to the GPT-3.5-turbo model and returns its response as a string.
// It returns an error if there are any issues during the process.
func ChatWithGPT(prompt string) (string, error) {
	var OpenAIAPIURL = config.Instance.Chatgpt.OpenAIAPIURL
	var model = config.Instance.Chatgpt.Model
	var apiKey = config.Instance.Chatgpt.OPENAI_API_KEY
	if apiKey == "" {
		return "", fmt.Errorf("error: OPENAI_API_KEY environment variable not set")
	}
	messages := []Message{
		{
			Role:    "user",
			Content: prompt,
		},
	}

	requestBody := &ChatCompletionRequest{Model: model, Messages: messages}
	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		return "", err
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", OpenAIAPIURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		response, err := io.ReadAll(resp.Body)
		return string(response), err
	}

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var response Response
	err = json.Unmarshal(responseBody, &response)
	if err != nil {
		return "", err
	}

	content := response.Choices[0].Message.Content
	return content, nil
}
