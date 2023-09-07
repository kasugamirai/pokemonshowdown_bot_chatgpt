package commands

import (
	"log"
	"strings"

	"xy.com/pokemonshowdownbot/bard"
	"xy.com/pokemonshowdownbot/chatgpt"
	"xy.com/pokemonshowdownbot/database"
	"xy.com/pokemonshowdownbot/models"
)

type HandlerFunc func(msg string) (string, error)

var CommandsMap map[string]HandlerFunc

func Prompt(msg string) (string, error) {
	response, err := chatgpt.ChatWithGPT(msg)
	if err != nil {
		log.Println("Error connect to ChatWithGPT:", err)
		return "", err
	}
	if len(response) < 218 {
		return response, nil
	} else {
		return "!code " + response, nil
	}
}

func Bard(msg string) (string, error) {
	response, err := bard.GenerateTextResponse(msg)
	if err != nil {
		log.Println("Error connect to bard:", err)
		return "", err
	}
	if len(response) < 218 {
		return response, nil
	} else {
		return "!code " + response, nil
	}
}

func AddSticker(msg string) (string, error) {
	msgs := strings.Split(msg, " ")
	if len(msgs) == 2 {
		_, err := models.AddSticker(database.DB, msgs[0], msgs[1])
		if err != nil {
			log.Println("Error connect to database:", err)
			return "", err
		}
	} else {
		return "Format error.", nil
	}
	return "The image has been successfully added to the database.", nil
}
func FindStickerByName(msg string) (string, error) {
	URL, err := models.FindStickerByName(database.DB, msg)
	if err != nil {
		log.Println(err)
		return "", err
	}
	return "/raw " + "<img src=\"" + URL.URL + "\" height=\"100\">" + "<br>", nil
}

func DeleteStickerBynameHander(msg string) (string, error) {
	msgs := strings.Split(msg, " ")
	if len(msgs) == 2 {
		_, err := models.DeleteStrikerByName(database.DB, msgs[1])
		if err != nil {
			log.Println("Error delete Sticker from database:", err)
			return "", err
		}
	} else {
		return "Format error.", nil
	}
	return "The sticker has been successfully delete from the database.", nil
}

func LoadCommands() {
	CommandsMap = map[string]HandlerFunc{
		"bard": Bard,
		"p":    Prompt,
		"add":  AddSticker,
		"find": FindStickerByName,
		"del":  DeleteStickerBynameHander,
	}
}
